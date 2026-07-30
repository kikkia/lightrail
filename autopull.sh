#!/bin/bash
set -e

ACTION=$1
REPO_DIR="/home/kikkia/lightrail"

export PATH=$PATH:/usr/local/go/bin:/home/kikkia/go/bin

export DISPLAY=:0
export XDG_RUNTIME_DIR="/run/user/$(id -u)"

if [ "$ACTION" = "sleep" ]; then
    echo "sleeping screen..."
    OUTPUT=$(xrandr | grep " connected" | cut -d' ' -f1 | head -n1)
    [ -n "$OUTPUT" ] && xrandr --output "$OUTPUT" --off

elif [ "$ACTION" = "wake" ]; then
    echo "waking screen..."
    OUTPUT=$(xrandr | grep " connected" | cut -d' ' -f1 | head -n1)
    [ -n "$OUTPUT" ] && xrandr --output "$OUTPUT" --auto
    
    echo "pulling master..."
    cd "$REPO_DIR"
    git checkout master
    git pull origin master
    
    echo "building binary..."
    go build -o lightrail .
    
    echo "restarting lightrail user service..."
    systemctl --user restart lightrail.service
fi