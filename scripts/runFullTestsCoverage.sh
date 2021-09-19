#!/bin/bash

go test ./... -coverprofile=coverage.out.tmp -v
cat coverage.out.tmp | grep -v -E ".obx.go|objectbox-model.go" > coverage.out
go tool cover -html=coverage.out
