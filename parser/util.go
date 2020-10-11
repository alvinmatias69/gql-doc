package parser

import (
	"fmt"
	"os"
	"path"

	"github.com/alvinmatias69/gql-doc/entity"
)

func isFileValid(searchPath string, docTypes []entity.DocType) error {
	for _, docType := range docTypes {
		curPath := path.Join(searchPath, string(docType))
		if _, err := os.Stat(curPath); os.IsNotExist(err) {
			return fmt.Errorf("Unable to locate %s", curPath)
		}
	}
	return nil
}
