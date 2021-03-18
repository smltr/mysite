#!/bin/bash
cd /home/ubuntu/mysite
go get ./cmd
mkdir bin
go build -o ./bin ./cmd