#!/bin/bash
set -e
go test -v
go build -o githubrepodispatch
env GOOS=linux GOARCH=amd64 go build -o githubrepodispatch_linux_amd64