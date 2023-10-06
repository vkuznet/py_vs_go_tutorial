### How to build and test Go data-service

Step 1: write your code (file `http_hello.go`):
```
package main

import (
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe(":8888", nil)
}
```

Bonus 1: let's write code with unusual syntax (file `h.go`):
```
package main

import (
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request){
                w.Write([]byte("hello world"))
                }

func main() {	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe(":8888", nil)}
```
and, now we can format it easily with go:
```
go fmt h.go
```

Bonus 2: let's get documentation about http:
```
go doc http
go doc http.Request
go doc http.ResponseWriter
go doc http.ListenAndServe
```

Step 2: build your code
```
go build -o http_hello

# if you need to build code for another architecture you'll do:
# for Linux AMD
GOOS=linux go build -o http_hello

# to specifically diable CGO support
CGO_ENABLED=0 GOOS=linux go build http_hello.go

# to strip extra (debuggin info) and reduce your executable size
GOOS=linux go build -ldflags="-s -w" http_hello.go

# be explicit with LDFLAGS
GOOS=linux go build -ldflags="-s -w -extldflags -static" http_hello.go

# to get idea about your environment please do
go env

# for Linux ARM
GOARCH=arm64 GOOS=linux go build -o http_hello

# for Linux Power8
GOARCH=ppc64le GOOS=linux go build

# for Windows AMD
GOARCH=amd64 GOOS=windows go build -o http_hello

# for Windows ARM
GOARCH=arm64 GOOS=windows go build -o http_hello

...
```

Step 3: run your service
```
./http_hello
```

Step 4: migrate your service to another node:
```
# either package your service or copye all codebase via ssh
scp http_hello vek3@lnx231.classe.cornell.edu:~/tutorial/go

# ssh to the node and run it
ssh lnx231
cd ~/tutorial/go
./http_hello
```

Step 5: test your service using
[curl](https://curl.se/)
and
[hey](https://github.com/rakyll/hey)
```
# use curl to perform HTTP request
curl http://localhost:8888

# use hey tool to perform stress tests
hey http://localhost:8888

# use 1k request and 100 concurrent clients:
hey -n 1000 -c 100 http://localhost:8888

# try harder with more concurrent clients
hey -n 1000 -c 200 http://localhost:8888
```

Step 6: work with 3rd party packages (`gin/server.go`):
```
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```
Let's try it out (`cd gin; vim server.go`):
```
# try to build it
go build

# setup new project
go mod init server
# at this point inspect go.mod
# or better way, if you'll put your code to github:
# go mod init github.com/vkuznet/server

# get all dependencies
go mod tidy
# at this point inspect go.mod and go.sum
```
And, now we can test it with curl/hey:
```
# inspect HTTP request flow
curl -v http://localhost:8080/ping

# perform stress test
hey http://localhost:8080/ping
```
