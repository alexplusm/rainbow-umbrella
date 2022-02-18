package repos

import (
	"rainbow-umbrella/internal/objects/dao"
)

func buildInsertOneUser(value *dao.User) *query {
	queryRaw := `
INSERT INTO users (
	login, hashed_password,
	first_name, last_name, birthday, gender, city,
	created_at
) VALUES (
	?, ?,
	?, ?, ?, ?, ?,
	?
);
`
	args := []interface{}{
		value.Login, value.HashedPassword,
		value.FirstName, value.LastName, value.Birthday, value.Gender, value.City,
		value.CreatedAt,
	}

	return &query{Query: queryRaw, Args: args}
}
