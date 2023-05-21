#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build .
zip -rq webhook.v1.0.linux.arm64.zip hook README.md