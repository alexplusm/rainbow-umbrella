package infrastruct

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"

	"rainbow-umbrella/internal/controllers"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/repos"
	"rainbow-umbrella/internal/services"
)

type IInjector interface {
	InjectUserController() interfaces.IUserController

	InjectSessionService() interfaces.ISessionService
}

type injector struct {
	dbClient    *sql.DB
	redisClient *redis.Client

	sessionService interfaces.ISessionService
}

func NewInjector(config *AppConfig) IInjector {
	dbClient, err := NewDBConn(config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	redisClient := NewRedisConn() // TODO: pass config end handle error

	sessionService := services.NewSessionService(repos.NewSessionRepo(redisClient))

	return &injector{dbClient: dbClient, redisClient: redisClient, sessionService: sessionService}
}

func (i injector) InjectUserController() interfaces.IUserController {
	return controllers.NewUserController(i.injectUserService())
}

func (i injector) InjectSessionService() interfaces.ISessionService {
	return i.sessionService
}

// ---

func (i injector) injectUserService() interfaces.IUserService {
	return services.NewUserService(i.injectUserRepo())
}

func (i injector) injectUserRepo() interfaces.IUserRepo {
	return repos.NewUserRepo(i.dbClient)
}

func (i injector) injectSessionRepo() interfaces.ISessionRepo {
	return repos.NewSessionRepo(i.redisClient)
}
