### Python server with dataframe
HTTP server with dataframes based on pandas package and Flask framework.

Here is benchmark numbers which is obtained from `run.sh` script
which provides both deployment procedures in each case and running the
server. The throughput is measured using
[hey](https://github.com/rakyll/hey) tool, e.g. `time ./run.sh`:
- total time to setup dependencies and start server: 17.5sec
- RSS: 25788, MEM 23MB, disk space: 221MB
- throughput: 559 req/sec
