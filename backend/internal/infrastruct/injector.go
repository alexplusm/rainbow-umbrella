package infrastruct

import (
	"database/sql"
	"log"

	"rainbow-umbrella/internal/controllers"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/repos"
	"rainbow-umbrella/internal/services"
)

type IInjector interface {
	InjectUserController() interfaces.IUserController
}

type injector struct {
	dbClient *sql.DB
}

func NewInjector() IInjector {
	dbClient, err := NewDBConn()
	if err != nil {
		log.Fatal(err)
	}

	return &injector{dbClient: dbClient}
}

func (i injector) InjectUserController() interfaces.IUserController {
	return controllers.NewUserController(i.injectUserService())
}

// ---

func (i injector) injectUserService() interfaces.IUserService {
	return services.NewUserService()
}

func (i injector) injectUserRepo() interfaces.IUserRepo {
	return repos.NewUserRepo(i.dbClient)
}
