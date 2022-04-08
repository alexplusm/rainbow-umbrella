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

	if err := utils.MakeDirs(consts.AppDirs); err != nil {
		log.Fatal(err)
	}

	injector := infrastruct.NewInjector(appConfig)
	userController := injector.InjectUserController()
	friendshipController := injector.InjectFriendshipController()

	r.Post("/api/v1/users/register", userController.Register)

	r.Get("/api/v1/users",
		NewSessionMiddleware(
			injector.InjectSessionService(), userController.List,
		),
	)

	r.Post("/api/v1/users/login", userController.Login)

	r.Get("/api/v1/users/{login}",
		NewSessionMiddleware(injector.InjectSessionService(), userController.Details),
	)

	// TODO: add NewSessionMiddleware
	r.Post("/api/v1/friendships", friendshipController.Create)
	r.Post("/api/v1/friendships/approve", friendshipController.Approve)
	r.Post("/api/v1/friendships/decline", friendshipController.Decline)
	r.Get("/api/v1/friendships/{login}", friendshipController.List)

	log.Printf("Start app on: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		log.Fatal(err)
	}
}

func NewLoggerMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logger!!!")
		f(w, r)
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

		userLogin, ok, err := sessionService.RetrieveUserLogin(sessionID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
				log.Printf("[NewSessionMiddleware][2]: %+v", err)
			}
			return
		}

		fmt.Println("@@@ [NewSessionMiddleware]: userLogin", userLogin)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(http.StatusText(http.StatusUnauthorized))); err != nil {
				log.Printf("[NewSessionMiddleware][3]: %+v", err)
			}
			return
		}

		ctx := sessionService.SetCurrentUserToCtx(r.Context(), userLogin)

		handler(w, r.WithContext(ctx))
	}
}
