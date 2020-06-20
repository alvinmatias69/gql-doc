package example

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestGenerator_Example(t *testing.T) {
	tests := []struct {
		name     string
		spec     entity.Spec
		template string
		want     entity.Spec
		wantErr  bool
	}{
		{
			name: "success generate simple query example",
			spec: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImageName",
						Type:       "String",
						IsScalar:   true,
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
			want: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImageName",
						Type:       "String",
						IsScalar:   true,
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
						Example: entity.Example{
							Request:  `{"id":1}`,
							Response: `{"getImageName":"Hello"}`,
						},
					},
				},
			},
		},
		{
			name: "success generate query with list",
			spec: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImageName",
						Type:       "String",
						IsScalar:   true,
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
					},
				},
			},
			want: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImageName",
						Type:       "String",
						IsScalar:   true,
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
						Example: entity.Example{
							Request:  `{"id":[1]}`,
							Response: `{"getImageName":"Hello"}`,
						},
					},
				},
			},
		},
		{
			name: "success generate query with user defined type",
			spec: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImages",
						Type:       "Image",
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
					},
				},
				Definitions: []entity.Definition{
					{
						Name: "Image",
						Properties: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
			want: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImages",
						Type:       "Image",
						IsNullable: true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
						Example: entity.Example{
							Request:  `{"id":[1]}`,
							Response: `{"getImages":{"id":1}}`,
						},
					},
				},
				Definitions: []entity.Definition{
					{
						Name: "Image",
						Properties: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success generate query with list of user defined type",
			spec: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImages",
						Type:       "Image",
						IsNullable: true,
						IsList:     true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
					},
				},
				Definitions: []entity.Definition{
					{
						Name: "Image",
						Properties: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
			want: entity.Spec{
				Queries: []entity.Property{
					{
						Name:       "getImages",
						Type:       "Image",
						IsNullable: true,
						IsList:     true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
						Example: entity.Example{
							Request:  `{"id":[1]}`,
							Response: `{"getImages":[{"id":1}]}`,
						},
					},
				},
				Definitions: []entity.Definition{
					{
						Name: "Image",
						Properties: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
		},
		{
			name: "success generate mutations with list user defined type",
			spec: entity.Spec{
				Mutations: []entity.Property{
					{
						Name:       "getImages",
						Type:       "Image",
						IsNullable: true,
						IsList:     true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
					},
				},
				Definitions: []entity.Definition{
					{
						Name: "Image",
						Properties: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
			want: entity.Spec{
				Mutations: []entity.Property{
					{
						Name:       "getImages",
						Type:       "Image",
						IsNullable: true,
						IsList:     true,
						Parameters: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
								IsList:   true,
							},
						},
						Example: entity.Example{
							Request:  `{"id":[1]}`,
							Response: `{"getImages":[{"id":1}]}`,
						},
					},
				},
				Definitions: []entity.Definition{
					{
						Name: "Image",
						Properties: []entity.Property{
							{
								Name:     "id",
								Type:     "Int",
								IsScalar: true,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.spec, tt.template)
			got, err := g.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.Example() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.Example() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
