package dao

type Friendship struct {
	ID               int
	RequestingUserID int
	TargetingUserID  int
	Status           string

	CreatedAt string
	UpdatedAt string
}
