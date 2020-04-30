#!/bin/bash
go build -ldflags "-s -w" -o bin/iris cmd/iris/iris.go