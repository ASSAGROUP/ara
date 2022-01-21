#!/bin/bash

# Build and run the example
bp await-port 5000
go run main.go build --context example localhost:5000/ara
go run main.go combine --context example localhost:5000/ara --all