#!/bin/bash

dir=$GOPATH

echo "GOPATH="$dir

go build -o gormdao ./cmd/gormdao/main.go

mv gormdao $dir/bin

echo "install success"