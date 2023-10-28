Basic tutorial to outline pros/cons of development simple data services using
python and go languages.

Please refer to specific topics below which explains how it is done in Go:
- [concurrency](go/concurrency/README.md)
- [memory](go/memory/README.md)
- [gin](go/gin/README.md)
- [dataframe](go/dataframe/README.md)
- [syntax](go/syntax/README.md)
- [tests](tests/README.md)

### hello world web server
Please see corresponding areas `py/http_py.py` for Python flask based 
web server and `go/http_hello.go` for Go (standard library) web
server.
- Python
```
# deployment procedure
python -m venv venv
source venv/bin/activat
pip install --upgrade pip
pip install flask
flask --app http_hello run

# to measure its performance use
hey http://localhost:5000
```
- GoLang
```
# start the server
go run http_hello.go

# to measure its performance use
hey http://localhost:8888
```

### dataframe web server
To benchmark each server we used `dataframe` within each server.
For Python it comes from `pandas` package while for Go
we rely on github.com/go-gota/gota/dataframe implementation

Here is benchmark numbers which is obtained from `run.sh` script
which provides both deployment procedures in each case and running the
server. The throughput is measured using
[hey](https://github.com/rakyll/hey) tool.
- Python (`cd py/dataframe && time ./run.sh`):
  - total time to setup dependencies and start server: 17.5sec
  - RSS: 25788, MEM 23MB, disk space: 221MB
  - throughput: 559 req/sec
- GoLang (`cd go/dataframe && time ./run.sh`):
  - total time to setup dependencies and start server: 1.3sec
  - RSS: 3660, MEM 1MB, disk space: 5MB
  - throughput: 12464 req/sec
