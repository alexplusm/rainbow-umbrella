package interfaces

import (
	"net/http"
)

type IUserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Details(w http.ResponseWriter, r *http.Request)
}
