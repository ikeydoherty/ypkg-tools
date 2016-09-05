#!/bin/bash

export GOPATH="$(pwd)"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOBIN"

(( stage_pass = 0 ))
(( stage_total = 0 ))

function stage(){
    (( pass = 0 ))
    (( total = 0 ))
    (( stage_total += 1 ))
    printf "\n\e[34m[%s]\e[39m\n" "$1"
}

function stage_end(){
    if [ $pass == $total ]; then
        (( stage_pass += 1 ))
        printf "\e[33m%-35s\e[32m%s\e[39m\n" "$1" "<PASS>"
    else
        printf "\e[33m%-35s\e[31m%s\e[39m\n" "$1" "<FAIL>"
    fi
}

function task(){
    printf "%-35s" "$1"
    (( total += 1 ))
}

function task_end(){
    if [ $? == 0 ]; then
        printf "\e[34m%s\e[39m\n" "<DONE>"
        (( pass += 1 ))
    else
        printf "\e[31m%s\e[39m\n" "<FAIL>"
    fi
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
    task "Installing golint..."
    go get -u github.com/golang/lint/golint
    rm -rf src/github.com/golang
    rm -rf src/golang.org
    rm -rf pkg
    task_end
fi

if [[ -e src/github.com/ikeydoherty/ypkg-tools ]]; then
    task "Removing existing build tree..."
	rm src/github.com/ikeydoherty/ypkg-tools -r
	task_end
fi

if [[ -e bin/yauto ]]; then
    task "Removing 'yauto' binary..."
    rm bin/yauto
    task_end
fi

task "Creating new build tree..."
mkdir -p src/github.com/ikeydoherty/ypkg-tools
ln -s "$(pwd)/ylib" src/github.com/ikeydoherty/ypkg-tools/.
ln -s "$(pwd)/ytools" src/github.com/ikeydoherty/ypkg-tools/.
task_end
stage_end "Setup stage finished:"


stage "Build"
task "Building 'yauto' binary..."
go install src/github.com/ikeydoherty/ypkg-tools/ytools/yauto/yauto.go
task_end
stage_end "Build stage finished:"


stage "Test"
go test ./...
stage_end "Testing stage finished:"

stage "Lint"
(( total += 1 ))
golint -set_exit_status ./...
if [ $? == 0 ]; then
    (( pass += 1 ))
fi
stage_end "Linting stage finished:"

stage "Cleanup"
if [[ -e src/github.com/ikeydoherty/ypkg-tools ]]; then
    task "Removing existing build tree..."
	rm src/github.com/ikeydoherty/ypkg-tools -r
	task_end
fi
stage_end "Cleanup stage finished:"

printf "\n"
if [ $stage_pass == $stage_total ]; then
    printf "\e[33m%-35s\e[32m%s\e[39m\n" "Build Completed:" "<PASS>"
else
    printf "\e[33m%-35s\e[31m%s\e[39m\n" "Build Completed:" "<FAIL>"
fi
printf "\n"