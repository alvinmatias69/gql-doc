package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
)

// Cli struct define apps configuration parsed from cli arguments
type Cli struct {
	input     *string
	output    *string
	template  *string
	types     *string
	quiet     *bool
	help      *bool
	noExample *bool
}

// New will init flag options then return cli struct
func New() *Cli {
	// init long arguments
	var (
		input     = flag.String("input", "./", "Input path")
		output    = flag.String("output", "", "Output path")
		template  = flag.String("template", string(entity.JSON), "Input path")
		types     = flag.String("type", "query,mutation", "Input path")
		quiet     = flag.Bool("quiet", false, "Input path")
		help      = flag.Bool("help", false, "Input path")
		noExample = flag.Bool("no-example", false, "do not generate query example")
	)

	// init shorthand arguments
	flag.StringVar(input, "i", "./", "Input path (shorthand)")
	flag.StringVar(output, "o", "", "Output path (shorthand)")
	flag.BoolVar(quiet, "q", false, "")
	flag.BoolVar(help, "h", false, "")

	return &Cli{
		input:     input,
		output:    output,
		template:  template,
		types:     types,
		quiet:     quiet,
		help:      help,
		noExample: noExample,
	}
}

// Parse argument flags then return it in entity.Config struct
func (c *Cli) Parse() entity.Config {
	flag.Parse()
	return entity.Config{
		Input:     *c.input,
		Output:    *c.output,
		Template:  *c.template,
		Type:      c.parseType(),
		Quiet:     *c.quiet,
		Help:      *c.help,
		NoExample: *c.noExample,
	}
}

// ShowHelp screen to stdout
func (c *Cli) ShowHelp() {
	var usage = `
Usage:
	gql-doc [--help | -h] [--quiet | -q] [--no-example] [--input | -i] [--output | -o] [--template] [--type]

Options:
	--help	 -h		Show help message
	--input	 -i		Set input search path [default: "./"]
	--output -o		Set output file, on empty will return to stdout [default: ""]
	--quiet	 -q		Set log verbosity to silent
	--template		Set template, if path provided it will use custom template ["json"|"markdown"|"html"|"confluence", default: "json"]
	--type			Set gql types to be generated separated by comma [default: "query,mutation"]
	--no-example    	Do not generate example for generated docs

Examples:
	- Generate docs from directory "./example/" and generate it to file "doc.json"
	  $ gql-doc -i example -o doc.json
	- Generate docs with custom template from "./custom-template.md"
	  $ gql-doc -i example -o doc.md --template custom-template.md
	`

	fmt.Println(usage)
}

func (c *Cli) parseType() []entity.DocType {
	var (
		typeSlice = strings.Split(*c.types, ",")
		result    = make([]entity.DocType, 0, len(typeSlice))
	)

	for _, typeData := range typeSlice {
		if data, ok := entity.DocTypeVariant[typeData]; ok {
			result = append(result, data)
		}
	}

	return result
}
