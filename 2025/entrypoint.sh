#!/bin/bash
cd 2025

if [[ $* == *"-v"* ]]; then
    echo "test mode verbose"
    packages=$(go list ./... 2>&1 | grep -v '/templateFolder')
    if [ $? -ne 0 ]; then
        echo "go list failed:"
        echo "$packages"
        exit 1
    fi
    go test $packages -v
elif [[ $* == *"-t"* ]]; then
    echo "test mode"
    packages=$(go list ./... 2>&1 | grep -v '/templateFolder')
    if [ $? -ne 0 ]; then
        echo "go list failed:"
        echo "$packages"
        exit 1
    fi
    go test $packages
else
    go run main.go $@
fi
