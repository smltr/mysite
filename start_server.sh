#!/bin/bash
echo "Launching server..."
tmux new -d -s mysite-session
tmux send -t mysite-session cd SPACE /home/ubuntu/mysite ENTER
tmux send -t mysite-session sudo SPACE ./server ENTER
echo "Server launched."