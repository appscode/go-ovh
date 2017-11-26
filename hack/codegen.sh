#! /bin/bash
set -xeou pipefail

GOPATH=$(go env GOPATH)
REPO_ROOT="$GOPATH/src/github.com/appscode/go-ovh"

pushd "$GOPATH/src/github.com/appscode/go-ovh"

rm -rf $REPO_ROOT/cloud/{client,models}; mkdir -p $REPO_ROOT/cloud; cd $REPO_ROOT/cloud
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A cloud --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

rm -rf $REPO_ROOT/domain/{client,models}; mkdir -p $REPO_ROOT/domain; cd $REPO_ROOT/domain
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A domain --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

rm -rf $REPO_ROOT/ip/{client,models}; mkdir -p $REPO_ROOT/ip; cd $REPO_ROOT/ip
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A ip --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

rm -rf $REPO_ROOT/iploadbalancing/{client,models}; mkdir -p $REPO_ROOT/iploadbalancing; cd $REPO_ROOT/iploadbalancing
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A ipLoadbalancing --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

rm -rf $REPO_ROOT/vip/{client,models}; mkdir -p $REPO_ROOT/vip; cd $REPO_ROOT/vip
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A vip --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

rm -rf $REPO_ROOT/vps/{client,models}; mkdir -p $REPO_ROOT/vps; cd $REPO_ROOT/vps
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A vps --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

rm -rf $REPO_ROOT/vrack/{client,models}; mkdir -p $REPO_ROOT/vrack; cd $REPO_ROOT/vrack
ovhapi2openapi -i api.yaml -o swagger.yaml
swagger generate client -f ./swagger.yaml -A vrack --copyright-file=$REPO_ROOT/hack/boilerplate.go.txt

popd
