#!/bin/bash

export GOPATH="$(pwd)"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOBIN"

function done_fail(){
    if [ $? == 0 ]; then
        printf "\e[34m%s\e[39m\n" "<DONE>"
    else
        printf "\e[31m%s\e[39m\n" "<FAIL>"
    fi
}

function pass_fail(){
    if [ $? == 0 ]; then
        printf "\e[32m%-35s\e[32m%s\e[39m\n" "$1" "<PASS>"
    else
        printf "\e[32m%-35s\e[31m%s\e[39m\n" "$1" "<FAIL>"
    fi
}

function stage(){
    printf "\n\e[34m[%s]\e[39m\n" "$1"
}

function status(){
    printf "%-35s" "$1"
}

printf "\n\e[34m%s\e[39m\n" "Running 'build.sh' for the 'ypkg-tools' project"

stage "Setup"
if [[ ! -e bin ]]; then
    mkdir bin
fi
if [[ ! -e src ]]; then
    mkdir src
fi
if [[ ! -e bin/golint ]]; then
    status "Installing golint..."
    go get -u github.com/golang/lint/golint
    rm -rf src/github.com/golang
    rm -rf src/golang.org
    rm -rf pkg
    done_fail
fi

if [[ -e src/github.com/ikeydoherty/ypkg-tools ]]; then
    status "Removing existing build tree..."
	rm src/github.com/ikeydoherty/ypkg-tools -r
	done_fail
fi

if [[ -e bin/yauto ]]; then
    status "Removing 'yauto' binary..."
    rm bin/yauto
    done_fail
fi
if [[ -e bin/ytools ]]; then
    status "Removing 'ytools' binary..."
    rm bin/ytools
    done_fail
fi


status "Creating new build tree..."
mkdir -p src/github.com/ikeydoherty/ypkg-tools
ln -s "$(pwd)/yauto.go" src/github.com/ikeydoherty/ypkg-tools/.
ln -s "$(pwd)/ylib" src/github.com/ikeydoherty/ypkg-tools/.
ln -s "$(pwd)/ytools" src/github.com/ikeydoherty/ypkg-tools/.
done_fail

stage "Build"
status "Building 'yauto' binary..."
go install src/github.com/ikeydoherty/ypkg-tools/yauto.go
done_fail

status "Building 'ytools' binary..."
go install src/github.com/ikeydoherty/ypkg-tools/ytools/ytools.go
done_fail

stage "Test"
go test ./...
pass_fail "Testing stage finished:"

stage "Lint"
golint -set_exit_status
pass_fail "Linting stage finished:"

printf "\n"