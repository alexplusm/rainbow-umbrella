package interfaces

import (
	"context"

	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type IUserRepo interface {
	InsertOne(item *dao.User) error
	List(filter *bo.UserFilter) ([]dao.User, error)
}

type ISessionRepo interface {
	InsertOne(sessionID, login string) error
	Exists(sessionID string) (bool, error)
	RetrieveUserLogin(sessionID string) (string, error)
}

type IFriendshipRepo interface {
	InsertOne(friendship *dao.Friendship) error
	FriendList(userID uint64) (*dao.FriendList, error)
	UpdateStatus(id uint64, status string) error
}

type IInterestRepo interface {
	InsertOne(ctx context.Context, value *dao.Interest) error
}
