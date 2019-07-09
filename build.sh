#!/bin/bash
echo "Building the binary ..."
GOOS=linux GOARCH=amd64 go build -o ipreporter ipreporter.go db.go
echo "Creating a deployment file ..."
zip deployment.zip ipreporter
echo "Cleaning up ..."
rm ipreporter
echo "Done."
