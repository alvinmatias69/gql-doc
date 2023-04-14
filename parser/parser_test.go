package parser

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/alvinmatias69/gql-doc/parser/extractor"
	"github.com/alvinmatias69/gql-doc/parser/matcher"
)

func TestParser_parse(t *testing.T) {
	type fields struct {
		searchPath string
	}
	type args struct {
		docType entity.DocType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Spec
		wantErr bool
	}{
		{
			name: "success parse package name",
			fields: fields{
				searchPath: "./test_data/package",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_package",
			},
		},
		{
			name: "success parse union",
			fields: fields{
				searchPath: "./test_data/union",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_union",
				Definitions: []entity.Definition{
					{
						Name:    "TestUnion",
						Variant: entity.Union,
						Comment: "this is comment",
						Properties: []entity.Property{
							{
								Name: "First",
							},
							{
								Name: "Second",
							},
						},
					},
				},
			},
		},
		{
			name: "success parse function",
			fields: fields{
				searchPath: "./test_data/function",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_function",
				Queries: []entity.Property{
					{
						Name:       "getProfile",
						Comment:    "get profile by name",
						Type:       "Profile",
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "name",
								Type:     "String",
								IsScalar: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success parse empty function",
			fields: fields{
				searchPath: "./test_data/function_empty",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_function",
				Queries: []entity.Property{
					{
						Name:       "getProfile",
						Comment:    "get profile by name",
						Type:       "Profile",
						IsNullable: true,
					},
				},
			},
		},
		{
			name: "success parse object",
			fields: fields{
				searchPath: "./test_data/object",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_object",
				Definitions: []entity.Definition{
					{
						Name:    "Response",
						Comment: "response data",
						Variant: entity.Object,
						Properties: []entity.Property{
							{
								Name:       "name",
								Type:       "String",
								IsNullable: true,
								IsScalar:   true,
							},
							{
								Name:     "age",
								Type:     "Int",
								IsScalar: true,
							},
							{
								Name:       "groupID",
								Type:       "Int",
								IsScalar:   true,
								IsList:     true,
								IsNullable: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success parse interface",
			fields: fields{
				searchPath: "./test_data/interface",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_interface",
				Definitions: []entity.Definition{
					{
						Name:    "Response",
						Comment: "response data",
						Variant: entity.Interface,
						Properties: []entity.Property{
							{
								Name:       "name",
								Type:       "String",
								IsNullable: true,
								IsScalar:   true,
							},
							{
								Name:     "age",
								Type:     "Int",
								IsScalar: true,
							},
							{
								Name:       "groupID",
								Type:       "Int",
								IsScalar:   true,
								IsList:     true,
								IsNullable: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success parse enum",
			fields: fields{
				searchPath: "./test_data/enum",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_enum",
				Definitions: []entity.Definition{
					{
						Name:    "Response",
						Comment: "response data",
						Variant: entity.Enum,
						Properties: []entity.Property{
							{
								Name: "SUCCESS",
							},
							{
								Name: "FAILURE",
							},
						},
					},
				},
			},
		},
		{
			name: "success parse complete",
			fields: fields{
				searchPath: "./test_data/complete",
			},
			args: args{
				docType: entity.Query,
			},
			want: entity.Spec{
				Name: "test_complete",
				Queries: []entity.Property{
					{
						Name:    "getProfile",
						Comment: "get profile by name",
						Type:    "Profile",
						Parameters: []entity.Property{
							{
								Name:     "name",
								Type:     "String",
								IsScalar: true,
							},
						},
						IsNullable: true,
					},
				},
				Definitions: []entity.Definition{
					{
						Name:    "Response",
						Comment: "response data",
						Variant: entity.Object,
						Properties: []entity.Property{
							{
								Name:       "name",
								Type:       "String",
								IsNullable: true,
								IsScalar:   true,
							},
							{
								Name:     "age",
								Type:     "Int",
								IsScalar: true,
							},
							{
								Name:       "groupID",
								Type:       "Int",
								IsNullable: true,
								IsScalar:   true,
								IsList:     true,
							},
						},
					},
					{
						Name:    "Input",
						Comment: "input data",
						Variant: entity.Input,
						Properties: []entity.Property{
							{
								Name:     "name",
								Type:     "String",
								IsScalar: true,
							},
						},
					},
					{
						Name:    "ResponseType",
						Comment: "response type variant",
						Variant: entity.Enum,
						Properties: []entity.Property{
							{
								Name: "SUCCESS",
							},
							{
								Name: "FAILURE",
							},
						},
					},
					{
						Name:    "SomeItf",
						Comment: "some interface",
						Variant: entity.Interface,
						Properties: []entity.Property{
							{
								Name:     "name",
								Type:     "String",
								IsScalar: true,
							},
						},
					},
					{
						Name:    "Dragon",
						Comment: "XYZ dragon union",
						Variant: entity.Union,
						Properties: []entity.Property{
							{
								Name: "X",
							},
							{
								Name: "Y",
							},
							{
								Name: "Z",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				searchPath: tt.fields.searchPath,
				matcher:    matcher.New(),
				extractor:  extractor.New(),
			}
			got, err := p.parse(tt.args.docType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parse() =\n%+v,\nwant\n%+v", got, tt.want)
			}
		})
	}
}
