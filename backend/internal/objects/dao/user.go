package dao

type User struct {
	ID             int64
	Login          string
	HashedPassword string

	FirstName string
	Lastname  string
	Birthday  string
	Gender    string
	City      string

	CreatedAt string
}
