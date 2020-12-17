#!/bin/bash

set -e

docker run --rm -it -v "$PWD":/go/src/github.com/foxever/go-xgboost -w /go/src/github.com/foxever/go-xgboost xgboost-testing:latest $@