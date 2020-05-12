package entity

type GQLDoc struct {
	Name     string `json:"name"`
	Query    *Doc   `json:"query,omitempty"`
	Mutation *Doc   `json:"mutation,omitempty"`
}
