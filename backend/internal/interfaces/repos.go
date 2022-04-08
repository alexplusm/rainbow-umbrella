package interfaces

import (
	"context"
	"database/sql"

	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type IUserRepo interface {
	InsertOne(ctx context.Context, item *dao.User) (uint64, error)
	List(filter *bo.UserFilter) ([]dao.User, error)
	RetrieveOne(ctx context.Context, login string) (*dao.User, error)
}

type ISessionRepo interface {
	InsertOne(sessionID, login string) error
	Exists(sessionID string) (bool, error)
	RetrieveUserLoginIfExist(ctx context.Context, sessionID string) (string, bool, error)
}

type IFriendshipRepo interface {
	InsertOne(friendship *dao.Friendship) error
	SelectOneByUsersLogin(ctx context.Context, login1, login2 string) (*dao.FriendshipWithLogin, error)
	FriendList(userID uint64) (*dao.FriendList, error)
	UpdateStatus(id uint64, status string) error
}

type IInterestRepo interface {
	InsertOne(ctx context.Context, value *dao.Interest) error
	InsertListAndAssignToUser(ctx context.Context, userID uint64, interests []string) error
	SelectListByUserID(tx *sql.Tx, ctx context.Context, userID uint64) ([]string, error)
}
