package parser

import (
	"os"
	"path/filepath"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/alvinmatias69/gql-doc/parser/extractor"
	"github.com/alvinmatias69/gql-doc/parser/matcher"
	"github.com/rs/zerolog/log"
)

type Parser struct {
	searchPath string
	docTypes   []entity.DocType
	matcher    *matcher.Matcher
	extractor  *extractor.Extractor
}

func New(inputPath string, docTypes []entity.DocType) (*Parser, error) {
	log.Info().Msg("Checking input file")
	searchPath := filepath.Clean(os.ExpandEnv(inputPath))
	err := isFileValid(searchPath, docTypes)

	if err != nil {
		return nil, err
	}

	return &Parser{
		searchPath: searchPath,
		docTypes:   docTypes,
		matcher:    matcher.New(),
		extractor:  extractor.New(),
	}, nil
}
