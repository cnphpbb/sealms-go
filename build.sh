#!/usr/bin/env sh


dirname=$(cd `dirname $0`; pwd)
startFile=$dirname/release/start.sh

cp -a $dirname/config.pro $dirname/release/
cp -a $dirname/public $dirname/release/
cp -a $dirname/swagger $dirname/release/
cp -a $dirname/template $dirname/release/

function linux()
{
    # linux x64
    GOPROXY="https://goproxy.io" GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/sealmsd
}

function mac()
{
    # mac X64
    GOPROXY="https://goproxy.io" CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o release/sealmsd
}

if [ $1 == "linux" ]; then
    linux
fi

if [ $1 == "mac" ]; then
    mac
fi