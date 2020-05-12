package templating

import (
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func Test_execToTemplate(t *testing.T) {
	type args struct {
		data     entity.GQLDoc
		tmplType entity.Template
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success return json string",
			args: args{
				data: entity.GQLDoc{
					Name: "Test",
					Query: &entity.Doc{
						Name: "gql docs",
						Functions: []entity.Function{
							{
								Name:       "parse",
								ReturnType: "String",
							},
						},
					},
				},
				tmplType: entity.JSON,
			},
			want: `{"name":"Test","query":{"name":"gql docs","functions":[{"name":"parse","parameters":null,"return_type":"String","comment":""}]}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToTemplate(tt.args.data, tt.args.tmplType)
			if (err != nil) != tt.wantErr {
				t.Errorf("execToTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("execToTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
