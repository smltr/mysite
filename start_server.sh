#!/bin/bash
tmux new -d -s mysite-session
tmux send -t mysite-session cd /home/ubuntu/mysite ENTER
tmux send -t mysite-session sudo SPACE ./server ENTER