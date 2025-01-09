#!/bin/bash
set -e
echo "Starting application..."
chmod +x /usr/local/bin/devops-playground
nohup /usr/local/bin/devops-playground > /var/log/devops-playground.log 2>&1 &
