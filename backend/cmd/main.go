package main

import (
	"fmt"
	"log"
	"net/http"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/infrastruct"
	"rainbow-umbrella/internal/utils"
)

func main() {
	port := "8080"

	if err := utils.MakeDirs(consts.AppDirs); err != nil {
		log.Fatal(err)
	}

	injector := infrastruct.NewInjector()
	userController := injector.InjectUserController()

	http.HandleFunc(
		"/api/v1/user/register",
		NewMethodMiddleware(http.MethodPost, userController.Register))

	log.Printf("Start app on: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
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
		handler(writer, request)
	}
}
