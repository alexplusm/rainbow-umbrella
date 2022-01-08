package interfaces

import (
	"net/http"
)

type IInjector interface {
}

type IUserController interface {
	Register(w http.ResponseWriter, r *http.Request)
}
