package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "<h1>Hello World!</h1>\nGreetings from %s with an %s CPU.", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3000", nil)
}
