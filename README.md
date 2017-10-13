# CI/CD with Go, Docker, Kubernetes, Codefresh, and Google Compute Engine

## 1.) Create an app

Let's create a Hello World! app in [Go](https://golang.org/) to keep things simple :D

```go
// main.go
package hello_go

import (
  "fmt"
  "net/http"
  "runtime"
)

func helloHandler(rw http.ResponseWriter, r *http.REquest) {
  fmt.Fprintf(rw, "<h1>Hello World!</h1>\nGreetings from %s with an %s CPU.", runtime.GOOS, runtime.GOARCH)
}

func main(){
  http.HandleFunc("/", helloHandler)
  http.ListenAndServe(":3000", nil)
}
```

## 2.) Create a [Docker](https://www.docker.com/) image  

Create the Dockerfile  

```Dockerfile
FROM golang

ADD . /go/src/github.com/ogryzek/hello_go
RUN go install github.com/ogryzek/hello_go
ENTRYPOINT /go/bin/hello_go

EXPOSE 3000
```

Build the image

```sh
docker build -t hello_go .
```

Run a container with the image we just built
```
docker run -p 3000:3000 hello_go
```

