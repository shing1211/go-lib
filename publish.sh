#!/bin/bash
go version
rm go.mod
rm go.sum
go mod init github.com/shing1211/go-lib
go mod tidy
set GOPROXY=proxy.golang.org
git tag v0.1.17
git push origin v0.1.17
go list -m github.com/shing1211/go-lib@v0.1.17
