package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/alvinmatias69/gql-doc/parser"
	"github.com/alvinmatias69/gql-doc/templating"
)

func main() {
	var inputPath = flag.String("ip", "./", "Input path")
	var outputPath = flag.String("out", "", "Output path")
	var template = flag.String("template", string(entity.JSON), "Template path")
	var isHelp = flag.Bool("help", false, "Show help")
	flag.Parse()

	if *isHelp {
		showUsage()
		return
	}

	inPath := resolvePath(*inputPath)

	query, err := parser.Parse(path.Join(inPath, "queries.go"))
	if err != nil {
		log.Println(err)
	}

	mutation, err := parser.Parse(path.Join(inPath, "mutations.go"))
	if err != nil {
		log.Println(err)
	}

	doc := entity.GQLDoc{
		Name:     query.Name,
		Query:    &query,
		Mutation: &mutation,
	}

	result, err := templating.ToTemplate(doc, entity.Template(*template))
	if err != nil {
		log.Println(err)
	}

	outPath := resolvePath(*outputPath)
	write(result, outPath)
}

func showUsage() {
	var usage = "\n" +
		"Usage:\n" +
		"\tgql-doc [-help] [-ip] [-out] [-template]" +
		"\n\n" +
		"Options:\n" +
		"\t-help\t\tShow help screen\n\n" +
		"\t-ip\t\tSet input search path, should include \n" +
		"\t\t\t`queries.go` and `mutations.go`\n" +
		"\t\t\t[ default: \"./\" ]\n\n" +
		"\t-out\t\tSet documentation output path\n" +
		"\t\t\ton empty the generated docs will be print to stdout\n" +
		"\t\t\t[ default: \"\" ]\n\n" +
		"\t-template\tUse template for documentation\n" +
		"\t\t\tsupply path to use custom template\n" +
		"\t\t\t[ html | markdown | default: json ]" +
		"\n\n" +
		"Example:\n" +
		"\tgqldoc -ip=affiliate -out=doc.json\n" +
		"\tgqldoc -ip=affiliate -out=doc.md -template=custom-md.tmpl\n"
	fmt.Println(usage)
}

func resolvePath(param string) string {
	param = os.ExpandEnv(param)
	return filepath.Clean(param)
}
