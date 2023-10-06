#!/bin/bash

# remove previous installation
rm -rf __pycache__ venv

# setup python environment and download dependencies
python -m venv venv
source venv/bin/activate
pip install flask
pip install pandas

# run http_df HTTP server
nohup flask --app http_df run 2>&1 1>& log < /dev/null &

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
curl -v http://localhost:5000/data
