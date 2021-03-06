#!/bin/bash
#If there is a terminal open running server, kill terminal
#then remove current site directory if it exists
cd /home/ubuntu
echo "Deleting site session..."
if tmux ls | grep -q mysite-session;
then
    tmux kill-ses -t mysite-session
    echo "Site session deleted."
else
    echo "Site session not found."
fi
echo "Deleting previous files..."
if [ -d mysite ]
then
    rmdir mysite
    echo "Site files deleted."
else
    echo "Site files not found."
fi