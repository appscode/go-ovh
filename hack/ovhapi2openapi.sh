#! /bin/bash
set -xeou pipefail

go get -u github.com/cygy/ovhapi2openapi

curl -o ovhapi.yaml https://raw.githubusercontent.com/cygy/ovhapi2openapi/master/Examples/OVHAPI-EU.yaml
ovhapi2openapi -i ovhapi.yaml -o cloud/swagger.yaml

swagger generate client -f ./cloud/swagger.yaml -A go-ovh
