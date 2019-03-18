#!/bin/bash
GOOS=linux go build
docker build -t cjzhang/tournaments .
go clean