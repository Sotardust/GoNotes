#!/bin/bash

function isCmdExist() {
	local cmd="$1"
  	if [ -z "$cmd" ]; then
		echo "Usage isCmdExist yourCmd"
		return 1
	fi

	which "$cmd" >/dev/null 2>&1
	if [ $? -eq 0 ]; then
		return 0
	fi

	return 2
}

function envFunc() {
    isCmdExist go
    if [ $? = "2" ] ;then
        echo "The Go environment is missing, please install Go@1.20"
        exit 0
    fi

    isCmdExist buf
    if [ $? = "2" ] ;then
        echo "The Buf Tool environment is missing and buf@v1.28.1 will be installed automatically"
        GO111MODULE=on go install github.com/bufbuild/buf/cmd/buf@v1.28.1
    fi

    isCmdExist buf
    if [ $? = "2" ] ;then
        echo "buf@v1.28.1 installation failed, please install manually"
        exit 0
    else
        echo "buf env ok"
    fi
}

case $1 in  
    env)  
        envFunc
        ;;  
    *)  
        echo "order invalid"  
        ;;  
esac