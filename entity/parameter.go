package entity

type Parameter struct {
	Name          string `json:"name"`
	ParamType     string `json:"param_type"`
	IsMandatory   bool   `json:"is_mandatory"`
	IsList        bool   `json:"is_list"`
	IsBuiltInType bool   `json:"is_built_in_type"`
	Comment       string `json:"comment"`
}
