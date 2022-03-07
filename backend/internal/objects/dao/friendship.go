package dao

type Friendship struct {
	ID               uint64
	RequestingUserID uint64
	TargetingUserID  uint64
	Status           string

	CreatedAt string
	UpdatedAt string
}
