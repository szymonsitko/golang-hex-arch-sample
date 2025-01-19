#!/bin/bash

# Test config must be an absolute pth
export TEST_CONFIG_FILE_PATH="/home/ssitko/Documents/projects/go_gin/.env"

go test -v ./...