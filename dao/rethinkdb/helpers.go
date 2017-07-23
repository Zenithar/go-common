package rethinkdb

import (
	"strings"

	r "gopkg.in/gorethink/gorethink.v3"

	"go.zenithar.org/common/dao/api"
)

// ConvertSortParameters to rethinkdb query string
func ConvertSortParameters(params api.SortParameters) []interface{} {

	var sorts []interface{}
	for k, v := range params {
		switch v {
		case api.Ascending:
			sorts = append(sorts, r.Asc(strings.ToLower(k)))
			break
		case api.Descending:
			sorts = append(sorts, r.Desc(strings.ToLower(k)))
			break
		default:
			sorts = append(sorts, r.Desc(strings.ToLower(k)))
		}
	}

	// Apply default sort
	if len(sorts) == 0 {
		sorts = append(sorts, r.Asc("id"))
	}

	return sorts
}
