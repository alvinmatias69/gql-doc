package entity

type Doc struct {
	Name      string     `json:"name"`
	Functions []Function `json:"functions,omitempty"`
	Types     []VarType  `json:"types,omitempty"`
	Enums     []Enum     `json:"enums,omitempty"`
}
