package interfaces

import (
	"context"

	"rainbow-umbrella/internal/objects/bo"
)

type IUserService interface {
	Register(user *bo.User) error
	RetrieveByLogin(login string) (*bo.User, error)
	LoginExist(login string) (bool, error)
	GenerateAvatarFileName(originalName string) string
	List(filter *bo.UserFilter) ([]bo.User, error)
	GetUsersFriendshipStatus(login1, login2 string) (string, error)
}

type ISessionService interface {
	Create(user *bo.User) (string, error)
	Exists(sessionID string) (bool, error)
	RetrieveUserLogin(sessionID string) (string, bool, error)

	SetCurrentUserToCtx(ctx context.Context, login string) context.Context
	GetCurrentUserFromCtx(ctx context.Context) (string, bool)
}

type IFriendshipService interface {
	Create(value *bo.Friendship) error
	FriendList(user *bo.User) (*bo.FriendList, error)
	UpdateStatus(id uint64, status string) error
	RetrieveByUsersID(userID1, userID2 uint64) (*bo.Friendship, error)
}

type IInterestService interface {
	CreateListForUser(ctx context.Context, userID uint64, interests []string) error
}
