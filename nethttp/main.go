package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// 2. /hello/:name -> my name is <name>
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		// 去掉前缀 "/hello/" 就拿到 name
		name := r.URL.Path[len("/hello/"):]
		fmt.Fprintf(w, "my name is %s", name)
	})

	// 启动服务
	http.ListenAndServe(":8080", nil)
}
