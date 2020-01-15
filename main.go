package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	var inputPath = flag.String("ip", "./", "Input path")
	var out = flag.String("out", "", "Output path")
	var template = flag.String("template", JSONType, "Template path")
	var isHelp = flag.Bool("help", false, "Show help")
	flag.Parse()

	if *isHelp {
		showUsage()
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	query, err := parse(path.Join(dir, "/", *inputPath, "queries.go"))
	if err != nil {
		log.Println(err)
	}

	mutation, err := parse(path.Join(dir, "/", *inputPath, "mutations.go"))
	if err != nil {
		log.Println(err)
	}

	doc := GQLDoc{
		Name:     query.Name,
		Query:    &query,
		Mutation: &mutation,
	}

	result, err := execToTemplate(doc, *template)
	if err != nil {
		log.Println(err)
	}

	write(result, *out)
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
