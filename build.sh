#!/bin/bash

printf "[Setup]\n"
printf "Checking for exsiting build tree...\n"
if [[ -e src ]]; then
    printf "Removing existing build tree..."
	rm src -r
	printf "DONE\n"
fi

printf "Creating new build tree..."
mkdir -p src/github.com/ikeydoherty/ypkg-tools
ln -s "$(pwd)/ylib" src/github.com/ikeydoherty/ypkg-tools/.
ln -s "$(pwd)/ytools" src/github.com/ikeydoherty/ypkg-tools/.
printf "DONE\n"

printf "\n[Build]\n"
GOPATH="$(pwd)" go build ./...

printf "\n[Vet]\n"
GOPATH="$(pwd)" go vet ./...

printf "\n[Test]\n"
GOPATH="$(pwd)" go test ./...
