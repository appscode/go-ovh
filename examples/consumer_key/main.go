package main

import (
	"fmt"

	"github.com/appscode/go-ovh"
)

func main() {
	// Create a client using credentials from config files or environment variables
	client, err := ovh.NewEndpointClient("ovh-ca")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	ckReq := client.NewCkRequest()

	// Allow GET method on /me
	ckReq.AddRules(ovh.ReadOnly, "/me")

	// Allow GET method on /xdsl and all its sub routes
	ckReq.AddRecursiveRules(ovh.ReadOnly, "/cloud")

	// Run the request
	response, err := ckReq.Do()
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	// Print the validation URL and the Consumer key
	fmt.Printf("Generated consumer key: %s\n", response.ConsumerKey)
	fmt.Printf("Please visit %s to validate it\n", response.ValidationURL)
}
