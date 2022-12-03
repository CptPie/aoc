#!/bin/bash
cd 2022 || exit

if [[ $* == *"-t"*  ]]; then
    echo "test mode"
    go test ./...
else
    go run main.go $@
fi