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
	Create(user *bo.User) (string, error)
	Exists(sessionID string) (bool, error)
}

type IFriendshipService interface {
}
