#!/bin/bash
set -e

ACTION=$1
REPO_DIR="/home/kikkia/lightrail"
export DISPLAY=:0

if [ "$ACTION" = "sleep" ]; then
    echo "sleeping screen..."
    xrandr --output $(xrandr | grep " connected" | cut -d' ' -f1 | head -n1) --off
elif [ "$ACTION" = "wake" ]; then
    echo "waking screen..."
    xrandr --output $(xrandr | grep " connected" | cut -d' ' -f1 | head -n1) --auto
    
    echo "pulling master"
    cd "$REPO_DIR"
    git pull origin master
    
    go build
    
    echo "restarting"
    sudo systemctl restart lightrail.service
fi