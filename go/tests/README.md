```
go mod init http
go mod tidy
go test -v .
go test -run Benchmark -bench=.
```
