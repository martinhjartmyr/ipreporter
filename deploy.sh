#!/bin/bash
echo "Deploying ..."
aws lambda update-function-code --function-name ipReporter --zip-file fileb://deployment.zip
echo "Cleaning up ..."
rm deployment.zip
echo "Done."
