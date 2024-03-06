#!/bin/sh

nginx -c  /etc/nginx/nginx.conf &

/app/fms-api -f /app/etc/fms.yaml
