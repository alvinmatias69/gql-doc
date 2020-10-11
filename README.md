# gql-doc
Simple tools for generating documentation from graphQL Specs
![gql-doc demo](resource/demo.gif)

## Installation

```sh
$ go get github.com/alvinmatias69/gql-doc
```

gql-doc project is developed using `go 1.13.5`. It's possible to use in another go version, but there's no guarantee that it will works as expected.

## Usage
```sh
Usage:
	gql-doc [--help | -h] [--quiet | -q] [--no-example] [--input | -i] [--output | -o] [--template] [--type]

Options:
	--help	 -h		Show help message
	--input	 -i		Set input search path [default: "./"]
	--output -o		Set output file, on empty will return to stdout [default: ""]
	--quiet	 -q		Set log verbosity to silent
	--template		Set template, if path provided it will use custom template ["json"|"markdown"|"html"|"confluence", default: "json"]
	--type			Set gql types to be generated separated by comma [default: "query,mutation"]
	--no-example    Do not generate example for generated docs

Examples:
	- Generate docs from directory "./example/" and generate it to file "doc.json"
	  $ gql-doc -i example -o doc.json
	- Generate docs with custom template from "./custom-template.md"
	  $ gql-doc -i example -o doc.md --template custom-template.md
```

For more example on specs file and generated output you can refer to [example](example_data) directory.

## GraphQL Specs

```go
package gqlDocumentation

var query = `
	getImages(id:[Int!]!): Images!
`

var queryType = `
# Query to get User Images
type Images {
	data: [String]
}

`
```

GraphQL specs is written in go. Generally used to generate general files of graphql project. Basically it has 3 main properties: package name, method, and types.
`gql-doc` will look for `queries.go` for gql query and `mutations.go` for gql mutations in given directory.

## Template
Templates are written in golang [template](https://golang.org/pkg/text/template/). You can provide a custom template for your needs. The given data is defined by `GQLDoc` struct at [entity.go](entity.go) file.

## TODO
- [ ] Tidy up functions and go comment
- [ ] Add ability to read custom query and mutation files
