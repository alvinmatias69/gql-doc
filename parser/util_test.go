package parser

import (
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func Test_isFileValid(t *testing.T) {
	type args struct {
		inputPath string
		docTypes  []entity.DocType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success return nil on available data",
			args: args{
				inputPath: "../example",
				docTypes: []entity.DocType{
					entity.Query,
				},
			},
		},
		{
			name: "failure on unavailable data",
			args: args{
				inputPath: "..",
				docTypes: []entity.DocType{
					entity.Query,
					entity.Mutation,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isFileValid(tt.args.inputPath, tt.args.docTypes); (err != nil) != tt.wantErr {
				t.Errorf("isFileValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
