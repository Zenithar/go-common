package api

import "strings"

// SortParameters contains a hashmap of field name with sort direction
type SortParameters map[string]SortDirection

// SortConverter convert a list of string to a SortParameters instance
func SortConverter(sorts []string) *SortParameters {
	params := SortParameters{}

	if sorts != nil && len(sorts) > 0 {
		for _, cond := range sorts {
			if len(strings.TrimSpace(cond)) > 0 {
				switch cond[0] {
				case '-':
					params[cond[1:]] = Descending
				case '+', ' ':
					params[cond[1:]] = Ascending
				default:
					params[cond] = Ascending
				}
			}
		}
	}

	return &params
}
