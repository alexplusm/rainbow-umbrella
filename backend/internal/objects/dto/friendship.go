package dto

type Friendship struct {
	RequestingUserID uint64 `json:"requestingUserId"`
	TargetingUserID  uint64 `json:"targetingUserId"`
}
