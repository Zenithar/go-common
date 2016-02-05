package unique

//go:generate mockery -name Generator -output mock

type Generator interface {
	Generate() string
}
