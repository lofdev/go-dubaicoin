#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="dubaicoin/cppjit-testrunner" docker-cppjit
docker build --tag="dubaicoin/python-testrunner" docker-python
docker build --tag="dubaicoin/go-testrunner" docker-go
