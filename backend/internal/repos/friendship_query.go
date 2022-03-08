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

func buildFriendListQuery(userID uint64) *query {
	queryRaw := `
SELECT
	f.friendship_id, f.requesting_user_id, f.targeting_user_id, f.status,
	u.user_id, u.login,
	u.first_name, u.last_name, u.birthday, u.gender, u.city,
	u.created_at
FROM friendships AS f
INNER JOIN users AS u
ON u.user_id = f.requesting_user_id OR u.user_id = f.targeting_user_id
WHERE 
	u.user_id != ?
	AND (f.requesting_user_id = ? OR f.targeting_user_id = ?) 
`
	args := []interface{}{userID, userID, userID}

	return &query{Query: queryRaw, Args: args}
}

func buildUpdateOneFriendshipQuery(id uint64, status string) *query {
	queryRaw := `
UPDATE friendships
SET status = ?
WHERE friendship_id = ?
;
`
	args := []interface{}{status, id}

	return &query{Query: queryRaw, Args: args}
}
