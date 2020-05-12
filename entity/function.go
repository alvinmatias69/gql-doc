package entity

type Function struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
	ReturnType string      `json:"return_type"`
	Comment    string      `json:"comment"`
}
