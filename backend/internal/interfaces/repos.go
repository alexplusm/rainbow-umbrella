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
}
