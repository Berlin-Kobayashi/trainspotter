#!/bin/sh

env GOARCH=mips env GOOS=linux go build $1
