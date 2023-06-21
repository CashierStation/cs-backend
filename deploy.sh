#!/usr/bin/env bash

# pull repo
git pull origin main
go build -o csbackend

./csbackend --migrate

# kill old process
killall csbackend
wait 3

# check if process is running, if not start while loop
if pgrep -f "csbackend" >/dev/null; then
    while (true) do
        echo "Restarting csbackend..."
        ./csbackend
        wait 3
        exitcode=$?
        echo "Exit code of command is $exitcode"
    done
else
    echo "Process is not running."
fi