#!/bin/bash
#Open terminal session and start server in the session
echo "Launching server..."
tmux new -d -s mysite-session
tmux send -t mysite-session cd SPACE /home/ubuntu/mysite ENTER
tmux send -t mysite-session sudo SPACE ./bin/cmd ENTER
echo "Server launched."