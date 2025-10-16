#!/bin/sh

raw_output=$(./bin/deanl ./tests/print_string.dl)

output=$(echo "$raw_output" | tr -d '\n\r')

if [ "$output" = "Hello, world!" ]; then 
    echo "Test passed"
else
    echo "Test failed"
fi