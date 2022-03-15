SRC = cmd/main.go
NAME = main

install:
	go mod download

run:
	go run ${SRC}

build:
	go build -o $(NAME) $(SRC)

gen_users:
	go run tests/data/index.go
