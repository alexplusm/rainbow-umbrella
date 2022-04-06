package repos

import (
	"rainbow-umbrella/internal/objects/dao"
)

func buildInsertOneInterestQuery(value *dao.Interest) *query {
	queryRaw := `
INSERT INTO interests (value)
VALUES (?)
;
`
	args := []interface{}{value}

	return &query{Query: queryRaw, Args: args}
}

func buildQ() *query {
	return &query{}
}
