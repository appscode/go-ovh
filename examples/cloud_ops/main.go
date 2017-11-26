package main

import (
	"fmt"

	"github.com/appscode/go-ovh"
	"strings"
)

func main() {
	// Create a client using credentials from config files or environment variables
	client, err := ovh.NewEndpointClient("ovh-ca")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	resp, err := client.Cloud().GetCloudProject(nil, client)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	fmt.Println(strings.Join(resp.Payload, ","))
}
