package entity

type Config struct {
	Input     string
	Output    string
	Template  string
	Type      []DocType
	Quiet     bool
	Help      bool
	NoExample bool
}

type DocType string

const (
	Query    = "queries.go"
	Mutation = "mutations.go"
)

var DocTypeVariant = map[string]DocType{
	"query":    Query,
	"mutation": Mutation,
}
