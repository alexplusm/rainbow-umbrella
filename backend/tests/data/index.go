package main

import (
	"fmt"
	"log"
	"time"

	"rainbow-umbrella/internal/infrastruct"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
	"rainbow-umbrella/internal/repos"
)

const UserCount = 3

func main() {
	appConfig, err := new(infrastruct.AppConfig).BuildFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := infrastruct.NewDBConn(appConfig.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repos.NewUserRepo(dbClient)

	userGeneratorInst := new(userGenerator).build()

	for i := 0; i < UserCount; i++ {
		user := userGeneratorInst.generateUserDAO()

		if err = userRepo.InsertOne(user); err != nil {
			log.Fatal(err)
		}
	}
}

type userGenerator struct {
	loginCounter   int64
	hashedPassword string
}

func (g *userGenerator) build() *userGenerator {
	g.loginCounter = 0
	g.hashedPassword = "$2a$10$F.QZ5yG1OVLeSjb.kNGKveVAt6oIT8IYi5bavieoel3A5PHDJlFNG"

	return g
}

func (g *userGenerator) generateLogin() string {
	g.loginCounter++
	return fmt.Sprintf("gen_login_%v", g.loginCounter)
}

func (g *userGenerator) generateUserDAO() *dao.User {
	userBO := new(bo.User)

	userBO.Login = g.generateLogin()
	userBO.HashedPassword = g.hashedPassword
	userBO.CreatedAt = time.Now()

	return new(dao.User).FromBO(userBO)
}
