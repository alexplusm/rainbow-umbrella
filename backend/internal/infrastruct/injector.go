package infrastruct

import (
	"database/sql"

	"rainbow-umbrella/internal/controllers"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/services"
)

type IInjector interface {
	InjectUserController() interfaces.IUserController
}

type injector struct {
	dbClient *sql.DB
}

func NewInjector() IInjector {
	dbClient, _ := NewDBConn()

	return &injector{dbClient: dbClient}
}

func (i injector) InjectUserController() interfaces.IUserController {
	return controllers.NewUserController(i.injectUserService())
}

// ---

func (i injector) injectUserService() interfaces.IUserService {
	return services.NewUserService()
}
