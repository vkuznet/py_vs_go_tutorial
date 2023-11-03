### Go server with dataframe
HTTP server with dataframes based on github.com/go-gota/gota/dataframe implementation.
Plese read and run `run.sh` script.

Here is benchmark numbers which is obtained from `run.sh` script
which provides both deployment procedures in each case and running the
server. The throughput is measured using
[hey](https://github.com/rakyll/hey) tool, e.g. `time ./run.sh`:
- total time to setup dependencies and start server: 1.3sec
- RSS: 3660, MEM 1MB, disk space: 5MB
- throughput: 12464 req/sec
