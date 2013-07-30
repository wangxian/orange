#!/bin/sh

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build .
mkdir -p bin/osx
mv orange bin/osx/orange

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
mkdir -p bin/linux
mv orange bin/linux/orange

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build .
mkdir -p bin/win/
mv orange.exe bin/win/orange.exe
