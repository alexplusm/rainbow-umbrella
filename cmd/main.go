package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.ListenAndServe(":80", )
	fmt.Println("RUN")

	http.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("one"))
	})

	http.HandleFunc("/two", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("two"))
	})

	http.ListenAndServe(":8080", nil)
}
