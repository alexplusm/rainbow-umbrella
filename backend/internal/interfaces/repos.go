package interfaces

import (
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
}

type IFriendshipRepo interface {
	InsertOne(friendship *dao.Friendship) error
}
