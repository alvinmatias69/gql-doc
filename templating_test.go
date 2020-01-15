package main

import "testing"

func Test_execToTemplate(t *testing.T) {
	type args struct {
		data     GQLDoc
		tmplType string
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
				data: GQLDoc{
					Name: "Test",
					Query: &Docs{
						Name: "gql docs",
						Functions: []Function{
							Function{
								Name:       "parse",
								ReturnType: "String",
							},
						},
					},
				},
				tmplType: JSONType,
			},
			want: `{"name":"Test","query":{"name":"gql docs","functions":[{"name":"parse","parameters":null,"return_type":"String","comment":""}]}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execToTemplate(tt.args.data, tt.args.tmplType)
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
