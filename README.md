[![GoDoc](https://godoc.org/github.com/appscode/go-ovh?status.svg)](http://godoc.org/github.com/appscode/go-ovh)
[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/go-ovh)](https://goreportcard.com/report/github.com/appscode/go-ovh)

# go-ovh
Golang client for OVH

## Motivation
The official [OVH API GO client](https://github.com/ovh/go-ovh) only handles credential creation and requests signing. This project generates idimatic GO client for OVH apis using [go-swagger](https://github.com/go-swagger/go-swagger) generator.


### Acknowledgement
 - [ovh/go-ovh](https://github.com/ovh/go-ovh): This project adapts official api wrappers from OVH
 - [cygy/ovhapi2openapi](https://github.com/cygy/ovhapi2openapi): A modified version is used to translate OVH API schema to Swagger 2.0 schema
 - [go-swagger/go-swagger](https://github.com/go-swagger/go-swagger): Used to generate GO client from Swagger 2.0 api schema
