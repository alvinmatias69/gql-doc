package main

type Parameter struct {
	Name          string `json:"name"`
	ParamType     string `json:"param_type"`
	IsMandatory   bool   `json:"is_mandatory"`
	IsList        bool   `json:"is_list"`
	IsBuiltInType bool   `json:"is_built_in_type"`
	Comment       string `json:"comment"`
}

type Function struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
	ReturnType string      `json:"return_type"`
	Comment    string      `json:"comment"`
}

type VarType struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
	Comment    string      `json:"comment"`
}

type Docs struct {
	Name      string     `json:"name"`
	Functions []Function `json:"functions,omitempty"`
	Types     []VarType  `json:"types,omitempty"`
}

type GQLDoc struct {
	Name     string `json:"name"`
	Query    *Docs  `json:"query,omitempty"`
	Mutation *Docs  `json:"mutation,omitempty"`
}
