package dao

type Friendship struct {
	ID               int64
	RequestingUserID int64
	TargetingUserID  int64
	Status           string

	CreatedAt string
	UpdatedAt string
}
