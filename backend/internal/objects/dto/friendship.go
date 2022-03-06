package dto

type Friendship struct {
	RequestingUserID int64 `json:"requestingUserID"`
	TargetingUserID  int64 `json:"targetingUserID"`
}
