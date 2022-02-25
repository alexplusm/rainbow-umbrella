package repos

import (
	"bytes"
	"fmt"
	"text/template"
)

type query struct {
	Query string
	Args  []interface{}
}

func applyFilterToQuery(query string, filter interface{}) (r string, err error) {
	t := template.New("applyFilterToQueryTemplate")
	buffer := new(bytes.Buffer)

	if t, err = t.Parse(query); err != nil {
		return "", fmt.Errorf("[gofnd][.ApplyFilterToQuery][1]: %+v", err)
	}

	if err = t.Execute(buffer, filter); err != nil {
		return "", fmt.Errorf("[gofnd][.ApplyFilterToQuery][2]: %+v", err)
	}

	return buffer.String(), nil
}
