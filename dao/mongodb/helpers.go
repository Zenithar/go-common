package mongodb

import (
	"fmt"
	"strings"

	"zenithar.org/go/common/dao/api"
)

// ConvertSortParameters to mongodb query string
func ConvertSortParameters(params api.SortParameters) []string {

	var sorts []string
	for k, v := range params {
		switch v {
		case api.Ascending:
			sorts = append(sorts, strings.ToLower(k))
			break
		case api.Descending:
			sorts = append(sorts, fmt.Sprintf("-%s", strings.ToLower(k)))
			break
		default:
			sorts = append(sorts, strings.ToLower(k))
		}
	}

	return sorts
}
