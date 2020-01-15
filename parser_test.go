package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Docs
		wantErr bool
	}{
		{
			name: "failure on file not found",
			args: args{
				path: "./test_data.go",
			},
			wantErr: true,
		},
		{
			name: "success read package name",
			args: args{
				path: "./test_data/package.go",
			},
			want: Docs{
				Name: "test_data",
			},
		},
		{
			name: "success read function",
			args: args{
				path: "./test_data/function.go",
			},
			want: Docs{
				Name: "test_data",
				Functions: []Function{
					Function{
						Name:       "getProduct",
						ReturnType: "Product",
						Parameters: []Parameter{
							Parameter{
								Name:          "limit",
								ParamType:     "Int",
								IsBuiltInType: true,
							},
							Parameter{
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
				path: "./test_data/var_type.go",
			},
			want: Docs{
				Name: "test_data",
				Types: []VarType{
					VarType{
						Name: "Response",
						Parameters: []Parameter{
							Parameter{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
							Parameter{
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
				path: "./test_data/fn_and_type.go",
			},
			want: Docs{
				Name: "test_data",
				Functions: []Function{
					Function{
						Name: "register",
						Parameters: []Parameter{
							Parameter{
								Name:          "name",
								ParamType:     "String",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
						},
						ReturnType: "Response",
					},
				},
				Types: []VarType{
					VarType{
						Name: "Response",
						Parameters: []Parameter{
							Parameter{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
							},
							Parameter{
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
				path: "./test_data/complete.go",
			},
			want: Docs{
				Name: "test_data",
				Functions: []Function{
					Function{
						Name: "register",
						Parameters: []Parameter{
							Parameter{
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
				Types: []VarType{
					VarType{
						Name:    "Response",
						Comment: "Response given after register new user",
						Parameters: []Parameter{
							Parameter{
								Name:          "success",
								ParamType:     "Boolean",
								IsBuiltInType: true,
								IsMandatory:   true,
								Comment:       "success indicate register result",
							},
							Parameter{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nparse() = %+v\nwant %+v\n", got, tt.want)
			}
		})
	}
}

func Test_extractParams(t *testing.T) {
	type args struct {
		strParams string
	}
	tests := []struct {
		name string
		args args
		want []Parameter
	}{
		{
			name: "success return nil on empty string",
			args: args{
				strParams: "",
			},
			want: nil,
		},
		{
			name: "success return one param",
			args: args{
				strParams: "(name: String)",
			},
			want: []Parameter{
				Parameter{
					Name:          "name",
					ParamType:     "String",
					IsBuiltInType: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractParams(tt.args.strParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractFunction(t *testing.T) {
	type args struct {
		strFn string
	}
	tests := []struct {
		name string
		args args
		want Function
	}{
		{
			name: "success return no params",
			args: args{
				strFn: "getProduct : Product!",
			},
			want: Function{
				Name:       "getProduct",
				ReturnType: "Product",
			},
		},
		{
			name: "success return one params",
			args: args{
				strFn: "getProduct (limit: Int!) : Product!",
			},
			want: Function{
				Name: "getProduct",
				Parameters: []Parameter{
					Parameter{
						Name:          "limit",
						ParamType:     "Int",
						IsMandatory:   true,
						IsBuiltInType: true,
					},
				},
				ReturnType: "Product",
			},
		},
		{
			name: "success return two params",
			args: args{
				strFn: "getProduct (limit: Int!, offset: Int!) : Product!",
			},
			want: Function{
				Name: "getProduct",
				Parameters: []Parameter{
					Parameter{
						Name:          "limit",
						ParamType:     "Int",
						IsMandatory:   true,
						IsBuiltInType: true,
					},
					Parameter{
						Name:          "offset",
						ParamType:     "Int",
						IsMandatory:   true,
						IsBuiltInType: true,
					},
				},
				ReturnType: "Product",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractFunction(tt.args.strFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseParams(t *testing.T) {
	type args struct {
		strParams string
	}
	tests := []struct {
		name string
		args args
		want Parameter
	}{
		{
			name: "success return name and params",
			args: args{
				strParams: "name: type",
			},
			want: Parameter{
				Name:      "name",
				ParamType: "type",
			},
		},
		{
			name: "success return built in status",
			args: args{
				strParams: "name: String",
			},
			want: Parameter{
				Name:          "name",
				ParamType:     "String",
				IsBuiltInType: true,
			},
		},
		{
			name: "success return mandatory status",
			args: args{
				strParams: "name: String!",
			},
			want: Parameter{
				Name:          "name",
				ParamType:     "String",
				IsBuiltInType: true,
				IsMandatory:   true,
			},
		},
		{
			name: "success return list status",
			args: args{
				strParams: "name: [String]!",
			},
			want: Parameter{
				Name:          "name",
				ParamType:     "String",
				IsBuiltInType: true,
				IsMandatory:   true,
				IsList:        true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseParams(tt.args.strParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractComment(t *testing.T) {
	type args struct {
		strComment string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return trailing slash on empty comment",
			args: args{
				strComment: "#",
			},
			want: "\n",
		},
		{
			name: "success remove multiple comment tag",
			args: args{
				strComment: "#######",
			},
			want: "\n",
		},
		{
			name: "success extract comment",
			args: args{
				strComment: "# this is comment",
			},
			want: "this is comment\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractComment(tt.args.strComment); got != tt.want {
				t.Errorf("extractComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
