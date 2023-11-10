#!/bin/bash

# remove previous installation
rm -f go.mod go.sum http_df

# start from scratch to build go project
go mod init dataframe
go mod tidy

# build go executable
go build -o http_df -ldflags="-s -w"

# run http_df server
nohup ./http_df 2>&1 1>& log < /dev/null &

# grab some statistics
ps auxww | grep http_df | grep -v grep
proc=`ps auxww | grep http_df | grep -v grep`
pid=`echo $proc | awk '{print $2}'`
echo
echo $proc

echo
echo "### memory usage based on ps -xm command"
ps -xm -o rss,comm -p $pid
echo
echo "### memory usage based on top command"
top -pid $pid -stats mem -l 1

echo
echo "total disk usage:"
du -ksh . 

echo
echo "check the service using curl"
curl -v http://localhost:8888/data

echo
echo "perform clean-up"
ps auxww | grep http_df | grep -v grep | awk '{print "kill -9 "$2""}' | /bin/sh
