#!/bin/bash
dir=$GOPATH
echo $dir

go build -o gormdao ./cmd/gormdao/main.go

mv gormdao $dir/bin

echo "install success"