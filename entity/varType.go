package entity

type VarType struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
	Comment    string      `json:"comment"`
}
