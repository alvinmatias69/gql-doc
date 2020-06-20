package parser

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/alvinmatias69/gql-doc/parser/extractor"
	"github.com/alvinmatias69/gql-doc/parser/matcher"
)

func TestNew(t *testing.T) {
	type args struct {
		inputPath string
		docTypes  []entity.DocType
	}
	tests := []struct {
		name    string
		args    args
		want    *Parser
		wantErr bool
	}{
		{
			name: "failure on error invalid file",
			args: args{
				inputPath: "..",
				docTypes: []entity.DocType{
					entity.Query,
					entity.Mutation,
				},
			},
			wantErr: true,
		},
		{
			name: "success init parser",
			args: args{
				inputPath: "../example",
				docTypes: []entity.DocType{
					entity.Query,
					entity.Mutation,
				},
			},
			want: &Parser{
				searchPath: "../example",
				docTypes: []entity.DocType{
					entity.Query,
					entity.Mutation,
				},
				matcher:   matcher.New(),
				extractor: extractor.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.inputPath, tt.args.docTypes)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
