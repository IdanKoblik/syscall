#!/bin/sh
cd /home/container

MODIFIED_STARTUP=$(echo "$STARTUP" | sed -e 's/{{/${/g' -e 's/}}/}/g')

echo -e "\033[1;33mlaunching: ${MODIFIED_STARTUP}\033[0m"
eval ${MODIFIED_STARTUP}

