SRC = cmd/main.go
NAME = main

install:
	go mod download

run:
	go run ${SRC}

build:
#	go build -mod=vendor -o $(NAME) $(SRC)
	go build -o $(NAME) $(SRC)

#publish:
#	sh scripts/publish.sh
