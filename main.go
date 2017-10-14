package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<h1>Hello World!</h1>
	
	<p>Greetings from %s with an %s CPU.
	
	<p>Check out the <a href="https://github.com/ogryzek/hello_go">source code</a>.
	`, runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3000", nil)
}
