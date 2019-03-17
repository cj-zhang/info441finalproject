#!/bin/bash
GOOS=linux go build
docker build -t cjzhang/smash .
docker build -t cjzhang/smashdb ../db
go clean