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
	InjectFriendshipController() interfaces.IFriendshipController

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

	redisClient, err := NewRedisConn(config.RedisConfig)
	if err != nil {
		log.Fatal(err)
	}

	sessionService := services.NewSessionService(repos.NewSessionRepo(redisClient))

	return &injector{dbClient: dbClient, redisClient: redisClient, sessionService: sessionService}
}

func (i injector) InjectUserController() interfaces.IUserController {
	return controllers.NewUserController(i.injectUserService(), i.sessionService)
}

func (i injector) InjectFriendshipController() interfaces.IFriendshipController {
	return controllers.NewFriendshipController(i.injectFriendshipService(), i.injectUserService())
}

func (i injector) InjectSessionService() interfaces.ISessionService {
	return i.sessionService
}

// ---

func (i injector) injectUserService() interfaces.IUserService {
	return services.NewUserService(i.injectUserRepo(), i.injectInterestService(), i.injectFriendshipService())
}

func (i injector) injectFriendshipService() interfaces.IFriendshipService {
	return services.NewFriendshipService(i.injectFriendshipRepo())
}

func (i injector) injectInterestService() interfaces.IInterestService {
	return services.NewInterestService(i.injectInterestRepo())
}

// --- repos

func (i injector) injectUserRepo() interfaces.IUserRepo {
	return repos.NewUserRepo(i.dbClient, i.injectInterestRepo())
}

func (i injector) injectSessionRepo() interfaces.ISessionRepo {
	return repos.NewSessionRepo(i.redisClient)
}

func (i injector) injectFriendshipRepo() interfaces.IFriendshipRepo {
	return repos.NewFriendshipRepo(i.dbClient)
}

func (i injector) injectInterestRepo() interfaces.IInterestRepo {
	return repos.NewInterestRepo(i.dbClient)
}
