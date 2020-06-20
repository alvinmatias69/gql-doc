package entity

type (
	TypeVariant   string
	ScalarVariant int
)

const (
	Scalar    = "Scalar"
	Object    = "Object"
	Interface = "Interface"
	Union     = "Union"
	Enum      = "Enum"
	Input     = "Input"
)

const (
	Int = iota
	Float
	String
	Boolean
	ID
)

type Definition struct {
	Name       string      `json:"name,omitempty"`
	Comment    string      `json:"comment,omitempty"`
	Variant    TypeVariant `json:"variant,omitempty"`
	Directive  string      `json:"directive,omitempty"`
	Properties []Property  `json:"properties,omitempty"`
}

type Property struct {
	Name       string     `json:"name,omitempty"`
	Comment    string     `json:"comment,omitempty"`
	Type       string     `json:"type,omitempty"`
	Directive  string     `json:"directive,omitempty"`
	Parameters []Property `json:"parameters,omitempty"`
	IsScalar   bool       `json:"is_scalar,omitempty"`
	IsNullable bool       `json:"is_nullable,omitempty"`
	IsList     bool       `json:"is_list,omitempty"`
	Example    Example    `json:"example,omitempty"`
}

type Example struct {
	Request  string `json:"request,omitempty"`
	Response string `json:"response,omitempty"`
}

type Spec struct {
	Name        string       `json:"name,omitempty"`
	Queries     []Property   `json:"queries,omitempty"`
	Mutations   []Property   `json:"mutations,omitempty"`
	Definitions []Definition `json:"definitions,omitempty"`
}

var ScalarTypes = map[string]ScalarVariant{
	"Int":     Int,
	"Float":   Float,
	"String":  String,
	"Boolean": Boolean,
	"ID":      ID,
}

var Types = map[string]TypeVariant{
	"type":      Object,
	"interface": Interface,
	"enum":      Enum,
	"input":     Input,
}
