#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0
export GOPATH=$ROOTDIR/go

#Get dependencies and test
go get github.com/smartystreets/goconvey/convey
go get ./...
go test ./...

#Compile code and export
cd accountservice;go get;go build -o ../devops/accountservice-linux-amd64;echo built `pwd`;cd ..
cd healthchecker;go get;go build -o ../devops/healthchecker-linux-amd64;echo built `pwd`;cd ..

#Copy binaries into output directory
cp devops/Dockerfile $ROOTDIR/image/.
cp devops/accountservice-linux-amd64 $ROOTDIR/image/.
cp devops/healthchecker-linux-amd64 $ROOTDIR/image/.
