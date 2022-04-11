package repos

import (
	"fmt"
	"github.com/alexplusm/gofnd"

	"rainbow-umbrella/internal/objects/bo"
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
)
;
`
	args := []interface{}{
		value.Login, value.HashedPassword,
		value.FirstName, value.LastName, value.Birthday, value.Gender, value.City,
		value.CreatedAt,
	}

	return &query{Query: queryRaw, Args: args}
}

func buildRetrieveOneUserQuery(login string) *query {
	queryRaw := `
SELECT
	user_id, login,
	first_name, last_name, birthday, gender, city
FROM users
	WHERE login = ?
;
`
	args := []interface{}{login}

	return &query{Query: queryRaw, Args: args}
}

func buildListUserQuery(filter *bo.UserFilter) (*query, error) {
	queryRaw := `
SELECT
	user_id, login, hashed_password,
	first_name, last_name, birthday, gender, city,
	created_at
FROM users
WHERE
	user_id > 0
	{{if .ByLogin}}
	AND login = ?
	{{end}}
	{{ if .ExcludeLogin }}
	AND login != ?
	{{ end }}
;
`
	args := make([]interface{}, 0)

	q, err := applyFilterToQuery(queryRaw, filter)
	if err != nil {
		return nil, fmt.Errorf("[buildListUserQuery][1]: %+v", err)
	}

	if filter.ByLogin != "" {
		args = append(args, filter.ByLogin)
	}

	if filter.ExcludeLogin != "" {
		args = append(args, filter.ExcludeLogin)
	}

	return &query{Query: q, Args: args}, nil
}

func buildListUserCommonInfoQuery(filter *bo.UserFilter) (*query, error) {
	queryRaw := `
SELECT 
	user_id, login, first_name, last_name
FROM users u
WHERE
	user_id > 0
{{- if not .ByLogins -}}
	AND (
		{{ range $index, $el := .ByLogins }}
			{{- if eq $index 0 -}}
			u.login = (?)
			{{- else -}}
			OR u.login = (?)
			{{- end -}}
		{{end -}}
		)
{{ end -}}
`
	queryStr, err := gofnd.ApplyFilterToQuery(queryRaw, filter)
	if err != nil {
		return nil, fmt.Errorf("[buildListUserCommonInfoQuery][1]: %w", err)
	}

	args := make([]interface{}, 0, 8)
	if len(filter.ByLogins) > 0 {
		args = append(args, filter.ByLoginsToInterface()...)
	}

	return &query{Query: queryStr, Args: args}, nil
}
