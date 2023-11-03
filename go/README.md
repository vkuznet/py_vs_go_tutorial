# Go tutorials
- [sytnax](syntax/README.md)
- [http](http/README.md)
- [tests](tests/README.md)
- [dataframe](dataframe/README.md)
- [memory](memory/README.md)
- [concurrency](concurrency/README.md)
- [gin framework](gin/README.md)
- [db](db/README.md)

### Installation instructions
```
# binary installation for Linux or other platforms:
curl -ksLO https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
tar xfz go1.21.3.linux-amd64.tar.gz
# setup go root location
export GOROOT=$PWD/go
# setup go path location (for locally installed packages)
export GOPATH=$PWD/gopath

# setup your PATH to get go executable
export PATH=$PATH:$GOROOT/bin

# you are ready to go
go version
go env
go doc http
```

For source installation please follow these
[instructions](https://go.dev/doc/install/source)
