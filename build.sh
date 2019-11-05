#!/bin/bash
set -e
go test -v
go build -o ghost2githubrepodispatch
env GOOS=linux GOARCH=amd64 go build -o ghost2githubrepodispatch_linux_amd64