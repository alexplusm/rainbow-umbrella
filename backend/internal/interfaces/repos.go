package interfaces

import (
	"rainbow-umbrella/internal/objects/dao"
)

type IUserRepo interface {
	InsertOne(item *dao.User) error
}

type ISessionRepo interface {
}
