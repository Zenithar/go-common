package api

// SortDirection is the enumeration for sort
type SortDirection int

const (
	// Ascending sort from bottom to up
	Ascending SortDirection = iota + 1
	// Descending sort from up to bottom
	Descending
)

var sortDirections = [...]string{
	"asc",
	"desc",
}

func (m SortDirection) String() string {
	return sortDirections[m-1]
}
