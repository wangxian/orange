#!/bin/sh

VERSION=3.5

mkdir -p bin/
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
mv orange bin/orange-osx-$VERSION

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
mv orange bin/orange-linux-$VERSION

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build
mv orange.exe bin/orange-win32-$VERSION.exe

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
mv orange.exe bin/orange-win64-$VERSION.exe