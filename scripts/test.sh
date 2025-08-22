#!/bin/sh

set -e

go run cmd/main.go -mode=debug -db=test.db -port=:8080