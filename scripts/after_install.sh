#!/bin/bash
#After files have been copied to directory
#create bin folder, build cmd folder and move
#binary to bin folder
cd /home/ubuntu/mysite
sudo mkdir bin
sudo go build -o ./bin ./cmd