#!/bin/bash
echo "Launching server..."
tmux new -d -s mysite-session
tmux send -t mysite-session cd SPACE /home/ubuntu/mysite/bin ENTER
tmux send -t mysite-session sudo SPACE ./cmd ENTER
echo "Server launched."