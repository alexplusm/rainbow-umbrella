package main

import (
	"fmt"
	"net/http"

	"rainbow-umbrella/internal/infrastruct"
)

func main() {
	fmt.Println("RUN")

	injector := infrastruct.NewInjector()
	injector.InjectUserController()

	http.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("one"))
	})

	http.HandleFunc("/two", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("two"))
	})

	http.ListenAndServe(":8080", nil)
}
