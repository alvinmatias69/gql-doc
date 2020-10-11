package entity

type Template string

const (
	JSON       Template = "json"
	Markdown   Template = "md"
	HTML       Template = "html"
	Confluence Template = "confluence"
)

var AvailableTemplate = map[string]Template{
	"json":       JSON,
	"markdown":   Markdown,
	"html":       HTML,
	"confluence": Confluence,
}
