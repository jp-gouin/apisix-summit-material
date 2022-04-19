#!/bin/sh
echo "10.199.102.15 myawesomeapi.pocpoc.poc" > /etc/hosts
exec nginx -g "daemon off;" &
/usr/share/api

