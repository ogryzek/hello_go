# CI/CD with Go, Docker, Kubernetes, Codefresh, and Google Compute Engine

## 1.) Create an app

Let's create a Hello World! app in Go to keep things simple :D

```go
// main.go
package main

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
