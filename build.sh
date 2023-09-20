#!/bin/bash

go clean --cache && go test -v -cover microservices/...
go build -o auth/authsvc auth/main.go
go build -o api/apisvc api/main.go