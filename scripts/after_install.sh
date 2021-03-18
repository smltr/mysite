#!/bin/bash
cd /home/ubuntu/mysite
go get ./cmd
sudo mkdir bin
go build -o ./bin ./cmd