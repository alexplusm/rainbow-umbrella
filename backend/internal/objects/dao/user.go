package dao

type User struct {
	ID             string
	Login          string
	HashedPassword string

	FirstName string
	Lastname  string
	Gender    string
	City      string
}
