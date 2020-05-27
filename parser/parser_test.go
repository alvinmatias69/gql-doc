package parser

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func Test_parse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Doc
		wantErr bool
	}{
		{
			name: "failure on file not found",
			args: args{
				path: "./test_data/hel",
			},
			wantErr: true,
		},
		{
			name: "success read package name",
			args: args{
				path: "./test_data/package",
			},
			want: entity.Doc{
				Name: "test_data",
			},
		},
		{
			name: "success read function",
			args: args{
				path: "./test_data/function",
			},
			want: entity.Doc{
				Name: "test_data",
				Functions: []entity.Function{
					{
						Name:       "getProduct",
						ReturnType: "Product",
						Parameters: []entity.Parameter{
							{
								Name:          "limit",
								ParamType:     "Int",
								IsBuiltInType: true,
							},
							{
								Name:          "nextCursor",
								ParamType:     "String",
								IsBuiltInType: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success read types",
			args: args{
				path: "./test_data/var_type",
			},
			want: entity.Doc{
				Name: "test_data",
				Types: []entity.VarType{
					{
						Name: "Response",
						Parameters: []entity.Parameter{
							{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
							{
								Name:          "error",
								ParamType:     "String",
								IsBuiltInType: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success read types and functions",
			args: args{
				path: "./test_data/fn_and_type",
			},
			want: entity.Doc{
				Name: "test_data",
				Functions: []entity.Function{
					{
						Name: "register",
						Parameters: []entity.Parameter{
							{
								Name:          "name",
								ParamType:     "String",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
						},
						ReturnType: "Response",
					},
				},
				Types: []entity.VarType{
					{
						Name: "Response",
						Parameters: []entity.Parameter{
							{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
							{
								Name:          "error",
								ParamType:     "String",
								IsBuiltInType: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success read comment",
			args: args{
				path: "./test_data/complete",
			},
			want: entity.Doc{
				Name: "test_data",
				Functions: []entity.Function{
					{
						Name: "register",
						Parameters: []entity.Parameter{
							{
								Name:          "name",
								ParamType:     "String",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
						},
						ReturnType: "Response",
						Comment:    "register new user",
					},
				},
				Types: []entity.VarType{
					{
						Name:    "Response",
						Comment: "Response given after register new user",
						Parameters: []entity.Parameter{
							{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
								Comment:       "success indicate register result",
							},
							{
								Name:          "error",
								ParamType:     "String",
								IsBuiltInType: true,
								Comment:       "error defines error message on failed register",
							},
						},
					},
				},
			},
		},
		{
			name: "success read input",
			args: args{
				path: "./test_data/input_type",
			},
			want: entity.Doc{
				Name: "test_data",
				Types: []entity.VarType{
					{
						Name: "Response",
						Parameters: []entity.Parameter{
							{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
							{
								Name:          "error",
								ParamType:     "String",
								IsBuiltInType: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success read enum",
			args: args{
				path: "./test_data/enum_type",
			},
			want: entity.Doc{
				Name: "test_data",
				Types: []entity.VarType{
					{
						Name: "Response",
						Parameters: []entity.Parameter{
							{
								Name:    "SUCCESS",
								Comment: "SUCCESS response",
							},
							{
								Name:    "FAILURE",
								Comment: "FAILURE response",
							},
						},
					},
				},
			},
		},
		{
			name: "success read enum with underscore",
			args: args{
				path: "./test_data/enum_with_underscore_type",
			},
			want: entity.Doc{
				Name: "test_data",
				Types: []entity.VarType{
					{
						Name: "Response",
						Parameters: []entity.Parameter{
							{
								Name:    "SUCCESS_RESPONSE",
								Comment: "SUCCESS response",
							},
							{
								Name:    "FAILURE_RESPONSE",
								Comment: "FAILURE response",
							},
						},
					},
				},
			},
		},
		{
			name: "success read union",
			args: args{
				path: "./test_data/union",
			},
			want: entity.Doc{
				Name: "test_data",
				Types: []entity.VarType{
					{
						Name:    "ImageQuery",
						Comment: "union ansemble",
						Parameters: []entity.Parameter{
							{
								Name: "min",
							},
							{
								Name: "meta",
							},
							{
								Name: "complete",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nparse() =\t%+v\nwant =\t\t%+v\n", got, tt.want)
			}
		})
	}
}
