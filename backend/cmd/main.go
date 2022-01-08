package main

import (
	"fmt"
	"log"
	"net/http"
	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/utils"

	"rainbow-umbrella/internal/infrastruct"
)

func main() {
	fmt.Println("RUN")

	if err := utils.MakeDirs(consts.AppDirs); err != nil {
		log.Fatal(err)
	}

	injector := infrastruct.NewInjector()
	userController := injector.InjectUserController()

	//router := http.NewServeMux()
	//router.HandleFunc()

	http.HandleFunc(
		"/api/v1/user/register",
		NewMethodMiddleware(http.MethodPost, userController.Register))

	http.HandleFunc("/two",
		NewLoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("two"))
		}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func NewLoggerMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("logger!!!")
		f(writer, request)
	}
}

func NewMethodMiddleware(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			if _, err := writer.Write([]byte(http.StatusText(http.StatusMethodNotAllowed))); err != nil {
				log.Println("[NewMethodMiddleware]: ", err.Error())
			}
		}
	}
}
