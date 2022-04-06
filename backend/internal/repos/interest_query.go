package repos

import (
	"fmt"

	"github.com/alexplusm/gofnd"

	"rainbow-umbrella/internal/objects/dao"
)

// TODO: remove
func buildInsertOneInterestQuery(value *dao.Interest) *query {
	queryRaw := `
INSERT INTO interests (value)
VALUES (?)
;
`
	args := []interface{}{value}

	return &query{Query: queryRaw, Args: args}
}

func buildInsertListInterestQuery(interests []string) (*query, error) {
	queryRaw := `
INSERT IGNORE INTO interests (value)
VALUES
{{range $index, $el := .Interests}}
	{{- if eq $index 0 -}}
	(?)
	{{- else -}}
	,(?)
	{{- end -}}
{{end -}}
;
`
	filter := struct{ Interests []string }{Interests: interests}
	queryStr, err := gofnd.ApplyFilterToQuery(queryRaw, &filter)
	if err != nil {
		return nil, fmt.Errorf("[buildInsertListInterestQuery][1]: %w", err)
	}

	args := make([]interface{}, len(interests))
	for i, value := range interests {
		args[i] = value
	}

	return &query{Query: queryStr, Args: args}, nil
}

func buildSelectInterestsIDsQuery(interests []string) (*query, error) {
	queryRaw := `
SELECT interest_id FROM interests
WHERE
{{range $index, $el := .Interests}}
	{{- if eq $index 0 -}}
	value = (?)
	{{- else -}}
	OR value = (?) 
	{{- end -}}
{{end -}}
;
`
	filter := struct{ Interests []string }{Interests: interests}
	queryStr, err := gofnd.ApplyFilterToQuery(queryRaw, &filter)
	if err != nil {
		return nil, fmt.Errorf("[buildSelectInterestsIDsQuery][1]: %w", err)
	}
	args := make([]interface{}, len(interests))
	for i, value := range interests {
		args[i] = value
	}

	return &query{Query: queryStr, Args: args}, nil
}
