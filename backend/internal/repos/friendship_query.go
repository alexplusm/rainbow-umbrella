package repos

import (
	"rainbow-umbrella/internal/objects/dao"
)

func buildInsertOneFriendshipQuery(value *dao.Friendship) *query {
	queryRaw := `
INSERT INTO friendships (
	requesting_user_id, targeting_user_id, status,
	created_at
) VALUES (
	?, ?, ?,
	?
);
`
	args := []interface{}{
		value.RequestingUserID, value.TargetingUserID, value.Status,
		value.CreatedAt,
	}

	return &query{Query: queryRaw, Args: args}
}
