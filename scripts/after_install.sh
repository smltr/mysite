#!/bin/bash
cd /home/ubuntu/mysite
sudo go get ./cmd
sudo mkdir bin
sudo go build -o ./bin ./cmd