package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"rainbow-umbrella/internal/infrastruct"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
	"rainbow-umbrella/internal/repos"
)

var (
	Cities = []string{
		"Paris", "Moscow", "New York", "London", "LosAngeles",
		"Berlin", "Piter", "Roma", "Oslo", "Amsterdam",
	}
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

	userGeneratorInst, err := new(userGenerator).build()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < UserCount; i++ {
		user := userGeneratorInst.generateUserDAO()

		if err = userRepo.InsertOne(user); err != nil {
			log.Fatal(err)
		}
	}
}

type userGenerator struct {
	loginCounter   int64
	loginPrefix    string
	hashedPassword string

	firstNames []string
	lastNames  []string
}

func (g *userGenerator) build() (*userGenerator, error) {
	rand.Seed(time.Now().Unix())

	g.loginCounter = 0
	g.hashedPassword = "$2a$10$F.QZ5yG1OVLeSjb.kNGKveVAt6oIT8IYi5bavieoel3A5PHDJlFNG"
	g.loginPrefix = fmt.Sprintf("gen_login_%v", time.Now().Unix())

	firstNamesRaw, err := ioutil.ReadFile("first_names.txt")
	if err != nil {
		return nil, err
	}
	for _, firstName := range strings.Split(string(firstNamesRaw), "\n") {
		if value := strings.TrimSpace(firstName); value != "" {
			g.firstNames = append(g.firstNames, value)
		}
	}

	lastNamesRaw, err := ioutil.ReadFile("last_names.txt")
	if err != nil {
		return nil, err
	}
	for _, lastName := range strings.Split(string(lastNamesRaw), "\n") {
		if value := strings.TrimSpace(lastName); value != "" {
			g.lastNames = append(g.lastNames, value)
		}
	}

	return g, nil
}

func (g *userGenerator) generateLogin() string {
	g.loginCounter++
	return fmt.Sprintf("%v_%v", g.loginPrefix, g.loginCounter)
}

func (g *userGenerator) generateCity() string {
	return Cities[rand.Intn(len(Cities))]
}

func (g *userGenerator) generateFirstName() string {
	return g.firstNames[rand.Intn(len(g.firstNames))]
}

func (g *userGenerator) generateLastName() string {
	return g.lastNames[rand.Intn(len(g.lastNames))]
}

func (g *userGenerator) generateGender() string {
	probability := rand.Float64()
	if probability >= 0.51 {
		return "Female"
	}
	return "Male"
}

func (g *userGenerator) generateUserDAO() *dao.User {
	userBO := new(bo.User)

	userBO.Login = g.generateLogin()
	userBO.HashedPassword = g.hashedPassword
	userBO.CreatedAt = time.Now()
	userBO.City = g.generateCity()
	userBO.FirstName = g.generateFirstName()
	userBO.LastName = g.generateLastName()
	userBO.Gender = g.generateGender()

	return new(dao.User).FromBO(userBO)
}
