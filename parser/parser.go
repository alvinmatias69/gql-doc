package parser

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/rs/zerolog/log"
)

func (p *Parser) Parse() (entity.Spec, error) {
	var (
		result = entity.Spec{}
		err    error
	)

	for _, docType := range p.docTypes {
		log.Info().Msgf("Parsing %s", docType)

		res, err := p.parse(docType)
		if err != nil {
			return result, fmt.Errorf("Error parsing %s. Message: %v", docType, err)
		}

		result.Name = res.Name
		result.Queries = append(result.Queries, res.Queries...)
		result.Mutations = append(result.Mutations, res.Mutations...)
		result.Definitions = append(result.Definitions, res.Definitions...)
	}

	return result, err
}

func (p *Parser) parse(docType entity.DocType) (entity.Spec, error) {
	file, err := os.Open(path.Join(p.searchPath, string(docType)))
	if err != nil {
		return entity.Spec{}, err
	}
	defer file.Close()

	var (
		name        string
		comment     string
		definitions []entity.Definition
		functions   []entity.Property
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		current = strings.TrimSpace(current)

		switch {
		case len(current) == 0:
			continue
		case p.matcher.Comment(current):
			comment += " " + p.extractor.Comment(current)
		case p.matcher.Package(current):
			name = p.extractor.Package(current)
		case p.matcher.Union(current):
			union, err := p.extractor.Union(current)
			if err != nil {
				return entity.Spec{}, fmt.Errorf(`error while processing "%s": %v`, current, err)
			}
			union.Comment = strings.TrimSpace(comment)
			definitions = append(definitions, union)
			comment = ""
		case p.matcher.Property(current):
			function, err := p.extractor.Property(current)
			if err != nil {
				return entity.Spec{}, fmt.Errorf(`error while processing "%s": %v`, current, err)
			}
			function.Comment = strings.TrimSpace(comment)
			functions = append(functions, function)
			comment = ""
		case p.matcher.Definition(current):
			definition, err := p.extractor.Definition(current)
			if err != nil {
				return entity.Spec{}, fmt.Errorf(`error while processing "%s": %v`, current, err)
			}
			definition.Comment = strings.TrimSpace(comment)
			comment = ""
			err = p.parseDefinition(&definition, scanner)
			if err != nil {
				return entity.Spec{}, err
			}
			definitions = append(definitions, definition)
		}
	}

	result := entity.Spec{
		Name:        name,
		Definitions: definitions,
	}

	if docType == entity.Query {
		result.Queries = functions
	} else {
		result.Mutations = functions
	}

	return result, nil
}

func (p *Parser) parseDefinition(definition *entity.Definition, scanner *bufio.Scanner) error {
	var comment string

	for scanner.Scan() {
		current := scanner.Text()
		current = strings.TrimSpace(current)

		switch {
		case len(current) == 0:
			continue
		case p.matcher.Comment(current):
			comment += " " + p.extractor.Comment(current)
		case p.matcher.Property(current):
			function, err := p.extractor.Property(current)
			if err != nil {
				return fmt.Errorf(`error while processing "%s": %v`, current, err)
			}
			function.Comment = strings.TrimSpace(comment)
			definition.Properties = append(definition.Properties, function)
			comment = ""
		case p.matcher.Enum(current):
			enum := p.extractor.Enum(current)
			enum.Comment = strings.TrimSpace(comment)
			definition.Properties = append(definition.Properties, enum)
			comment = ""
		case current == "}":
			return nil
		}
	}

	return nil
}
