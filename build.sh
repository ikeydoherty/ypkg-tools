#!/bin/bash

if [[ -e src ]]; then
	rm src -r
fi

mkdir -p src/github.com/ikeydoherty/ypkg-tools
ln -s "$(pwd)/ylib" src/github.com/ikeydoherty/ypkg-tools/.

GOPATH="$(pwd)" go build yauto.go

pushd ylib
GOPATH="$(pwd)" go test .
