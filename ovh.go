// Package ovh provides a HTTP wrapper for the OVH API.
package ovh

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	cloud_ops "github.com/appscode/go-ovh/cloud/client/operations"
	domain_ops "github.com/appscode/go-ovh/domain/client/operations"
	ip_ops "github.com/appscode/go-ovh/ip/client/operations"
	iploadbalancing_ops "github.com/appscode/go-ovh/iploadbalancing/client/operations"
	vip_ops "github.com/appscode/go-ovh/vip/client/operations"
	vps_ops "github.com/appscode/go-ovh/vps/client/operations"
	vrack_ops "github.com/appscode/go-ovh/vrack/client/operations"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// DefaultTimeout api requests after 180s
const DefaultTimeout = 180 * time.Second

// Endpoints
const (
	OvhEU        = "https://eu.api.ovh.com/1.0"
	OvhCA        = "https://ca.api.ovh.com/1.0"
	KimsufiEU    = "https://eu.api.kimsufi.com/1.0"
	KimsufiCA    = "https://ca.api.kimsufi.com/1.0"
	SoyoustartEU = "https://eu.api.soyoustart.com/1.0"
	SoyoustartCA = "https://ca.api.soyoustart.com/1.0"
	RunaboveCA   = "https://api.runabove.com/1.0"
)

// Endpoints conveniently maps endpoints names to their URI for external configuration
var Endpoints = map[string]string{
	"ovh-eu":        OvhEU,
	"ovh-ca":        OvhCA,
	"kimsufi-eu":    KimsufiEU,
	"kimsufi-ca":    KimsufiCA,
	"soyoustart-eu": SoyoustartEU,
	"soyoustart-ca": SoyoustartCA,
	"runabove-ca":   RunaboveCA,
}

// Errors
var (
	ErrAPIDown = errors.New("go-vh: the OVH API is down, it does't respond to /time anymore")
)

// Client represents a client to call the OVH API
type Client struct {
	// Self generated tokens. Create one by visiting
	// https://eu.api.ovh.com/createApp/
	// AppKey holds the Application key
	AppKey string

	// AppSecret holds the Application secret key
	AppSecret string

	// ConsumerKey holds the user/app specific token. It must have been validated before use.
	ConsumerKey string

	// API endpoint
	endpoint string

	// Client is the underlying HTTP client used to run the requests. It may be overloaded but a default one is instanciated in ``NewClient`` by default.
	Client *http.Client

	// Ensures that the timeDelta function is only ran once
	// sync.Once would consider init done, even in case of error
	// hence a good old flag
	timeDeltaMutex *sync.Mutex
	timeDeltaDone  bool
	timeDelta      time.Duration
	Timeout        time.Duration

	cloud           *cloud_ops.Client
	domain          *domain_ops.Client
	ip              *ip_ops.Client
	iploadbalancing *iploadbalancing_ops.Client
	vip             *vip_ops.Client
	vps             *vps_ops.Client
	vrack           *vrack_ops.Client

	Transport *httptransport.Runtime
}

// NewClient represents a new client to call the API
func NewClient(endpoint, appKey, appSecret, consumerKey string) (*Client, error) {
	client := Client{
		AppKey:         appKey,
		AppSecret:      appSecret,
		ConsumerKey:    consumerKey,
		Client:         &http.Client{},
		timeDeltaMutex: &sync.Mutex{},
		timeDeltaDone:  false,
		Timeout:        time.Duration(DefaultTimeout),
	}

	// Get and check the configuration
	if err := client.loadConfig(endpoint); err != nil {
		return nil, err
	}

	// create transport and client
	u, _ := url.Parse(client.endpoint)
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
		req.SetHeaderParam("X-Ovh-Application", client.AppKey)
		req.SetTimeout(client.Timeout)
		return nil
	})
	// transport.Debug = true
	client.cloud = cloud_ops.New(transport, strfmt.Default)
	client.domain = domain_ops.New(transport, strfmt.Default)
	client.ip = ip_ops.New(transport, strfmt.Default)
	client.iploadbalancing = iploadbalancing_ops.New(transport, strfmt.Default)
	client.vip = vip_ops.New(transport, strfmt.Default)
	client.vps = vps_ops.New(transport, strfmt.Default)
	client.vrack = vrack_ops.New(transport, strfmt.Default)

	client.Transport = transport
	return &client, nil
}

// NewEndpointClient will create an API client for specified
// endpoint and load all credentials from environment or
// configuration files
func NewEndpointClient(endpoint string) (*Client, error) {
	return NewClient(endpoint, "", "", "")
}

// NewDefaultClient will load all it's parameter from environment
// or configuration files
func NewDefaultClient() (*Client, error) {
	return NewClient("", "", "", "")
}

//
// High level helpers
//

// Ping performs a ping to OVH API.
// In fact, ping is just a /auth/time call, in order to check if API is up.
func (c *Client) Ping() error {
	_, err := c.getTime()
	return err
}

// TimeDelta represents the delay between the machine that runs the code and the
// OVH API. The delay shouldn't change, let's do it only once.
func (c *Client) TimeDelta() (time.Duration, error) {
	return c.getTimeDelta()
}

// Time returns time from the OVH API, by asking GET /auth/time.
func (c *Client) Time() (*time.Time, error) {
	return c.getTime()
}

//
// Low level API access
//

// getResult check the response and unmarshals it into the response type if needed.
// Helper function, called from CallAPI.
func (c *Client) getResponse(response *http.Response, resType interface{}) error {
	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// < 200 && >= 300 : API error
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		apiError := &APIError{Code: response.StatusCode}
		if err = json.Unmarshal(body, apiError); err != nil {
			apiError.Message = string(body)
		}
		apiError.QueryID = response.Header.Get("X-Ovh-QueryID")

		return apiError
	}

	// Nothing to unmarshal
	if len(body) == 0 || resType == nil {
		return nil
	}

	return json.Unmarshal(body, &resType)
}

// timeDelta returns the time  delta between the host and the remote API
func (c *Client) getTimeDelta() (time.Duration, error) {

	if !c.timeDeltaDone {
		// Ensure only one thread is updating
		c.timeDeltaMutex.Lock()

		// Did we wait ? Maybe no more needed
		if !c.timeDeltaDone {
			ovhTime, err := c.getTime()
			if err != nil {
				return 0, err
			}

			c.timeDelta = time.Since(*ovhTime)
			c.timeDeltaDone = true
		}
		c.timeDeltaMutex.Unlock()
	}

	return c.timeDelta, nil
}

// getTime t returns time from for a given api client endpoint
func (c *Client) getTime() (*time.Time, error) {
	var timestamp int64

	err := c.callAPI(http.MethodGet, "/auth/time", nil, &timestamp)
	if err != nil {
		return nil, err
	}

	serverTime := time.Unix(timestamp, 0)
	return &serverTime, nil
}

// getLocalTime is a function to be overwritten during the tests, it return the time
// on the the local machine
var getLocalTime = func() time.Time {
	return time.Now()
}

// getEndpointForSignature is a function to be overwritten during the tests, it returns a
// the endpoint
var getEndpointForSignature = func(c *Client) string {
	return c.endpoint
}

// CallAPI is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everyrthing went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *Client) callAPI(method, path string, reqBody, resType interface{}) error {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return err
		}
	}

	target := fmt.Sprintf("%s%s", c.endpoint, path)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return err
	}

	// Inject headers
	if body != nil {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("X-Ovh-Application", c.AppKey)
	req.Header.Add("Accept", "application/json")

	// Send the request with requested timeout
	c.Client.Timeout = c.Timeout
	response, err := c.Client.Do(req)

	if err != nil {
		return err
	}

	// Unmarshal the result into the resType if possible
	return c.getResponse(response, resType)
}

// AuthenticateRequest adds authentication data to the request
func (c *Client) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	timeDelta, err := c.TimeDelta()
	if err != nil {
		return err
	}

	timestamp := getLocalTime().Add(-timeDelta).Unix()

	req.SetHeaderParam("X-Ovh-Application", c.AppKey)
	req.SetHeaderParam("X-Ovh-Timestamp", strconv.FormatInt(timestamp, 10))
	req.SetHeaderParam("X-Ovh-Consumer", c.ConsumerKey)

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s+%s+%s+%s%s+%s+%d",
		c.AppSecret,
		c.ConsumerKey,
		req.GetMethod(),
		getEndpointForSignature(c),
		req.GetPath(),
		string(req.GetBody()),
		timestamp,
	)))
	req.SetHeaderParam("X-Ovh-Signature", fmt.Sprintf("$1$%x", h.Sum(nil)))
	req.SetTimeout(c.Timeout)
	return nil
}

//
// API services
//

func (c *Client) Cloud() *cloud_ops.Client {
	return c.cloud
}

func (c *Client) Domain() *domain_ops.Client {
	return c.domain
}

func (c *Client) IP() *ip_ops.Client {
	return c.ip
}

func (c *Client) IPLoadbalancing() *iploadbalancing_ops.Client {
	return c.iploadbalancing
}

func (c *Client) VIP() *vip_ops.Client {
	return c.vip
}

func (c *Client) VPS() *vps_ops.Client {
	return c.vps
}

func (c *Client) VRack() *vrack_ops.Client {
	return c.vrack
}
