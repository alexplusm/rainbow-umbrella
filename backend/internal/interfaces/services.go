package interfaces

import (
	"rainbow-umbrella/internal/objects/bo"
)

type IUserService interface {
	Register(user *bo.User) error
	RetrieveByLogin(login string) (*bo.User, error)
	LoginExist(login string) (bool, error)
	GenerateAvatarFileName(originalName string) string
	List(filter *bo.UserFilter) ([]bo.User, error)
}

type ISessionService interface {
	Create(user *bo.User) (string, error)
	Exists(sessionID string) (bool, error)
	RetrieveUserLogin(sessionID string) (string, bool, error)
}

type IFriendshipService interface {
	Create(value *bo.Friendship) error
	FriendList(user *bo.User) (*bo.FriendList, error)
	UpdateStatus(id uint64, status string) error
}

type IInterestService interface {
}
