package interfaces

import (
	"rainbow-umbrella/internal/objects/bo"
)

type IUserService interface {
	Register(user *bo.User) error
	RetrieveByLogin(login string) (*bo.User, error)
	LoginExist(login string) (bool, error)
	GenerateAvatarFileName(originalName string) string
}

type ISessionService interface {
}
