package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/infrastruct"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/utils"
)

func main() {
	port := "8080"

	r := chi.NewRouter()

	appConfig, err := new(infrastruct.AppConfig).BuildFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	infrastruct.NewDBConn(appConfig.DatabaseConfig)

	if err := utils.MakeDirs(consts.AppDirs); err != nil {
		log.Fatal(err)
	}

	injector := infrastruct.NewInjector(appConfig)
	userController := injector.InjectUserController()

	r.Post("/api/v1/user/register",
		NewMethodMiddleware(http.MethodPost, userController.Register))

	r.Post("/api/v1/user/login",
		NewMethodMiddleware(http.MethodPost, userController.Login))

	r.Get("/api/v1/user/{login}", NewSessionMiddleware(
		injector.InjectSessionService(),
		NewMethodMiddleware(http.MethodGet, userController.Details)))

	log.Printf("Start app on: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
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

func NewSessionMiddleware(sessionService interfaces.ISessionService, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get("X-SessionId")
		if sessionID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(http.StatusText(http.StatusUnauthorized))); err != nil {
				log.Printf("[NewSessionMiddleware][1]: %+v", err)
			}
			return
		}

		ok, err := sessionService.Exists(sessionID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
				log.Printf("[NewSessionMiddleware][2]: %+v", err)
			}
			return
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(http.StatusText(http.StatusUnauthorized))); err != nil {
				log.Printf("[NewSessionMiddleware][3]: %+v", err)
			}
			return
		}

		handler(w, r)
	}
}
