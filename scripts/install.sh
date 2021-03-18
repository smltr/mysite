#!/bin/bash
cd /home/ubuntu/mysite
go get ./cmd
go build -o ./bin ./cmd