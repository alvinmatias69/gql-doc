package extractor

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestExtractor_Property(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		e       *Extractor
		args    args
		want    entity.Property
		wantErr bool
	}{
		{
			name: "failure on error invalid property",
			args: args{
				input: "getProperty",
			},
			wantErr: true,
		},
		{
			name: "success on named type property",
			args: args{
				input: "getProperty : property",
			},
			want: entity.Property{
				Name:       "getProperty",
				Type:       "property",
				IsNullable: true,
			},
		},
		{
			name: "success on scalar type property",
			args: args{
				input: "getProperty : Int",
			},
			want: entity.Property{
				Name:       "getProperty",
				Type:       "Int",
				IsScalar:   true,
				IsNullable: true,
			},
		},
		{
			name: "success on scalar non-nullable type property",
			args: args{
				input: "getProperty : Int!",
			},
			want: entity.Property{
				Name:       "getProperty",
				Type:       "Int",
				IsScalar:   true,
				IsNullable: false,
			},
		},
		{
			name: "success on scalar non-nullable list type property",
			args: args{
				input: "getProperty : [Int]!",
			},
			want: entity.Property{
				Name:       "getProperty",
				Type:       "Int",
				IsScalar:   true,
				IsNullable: false,
				IsList:     true,
			},
		},
		{
			name: "success on property with params",
			args: args{
				input: "getProperty(name: String) : Int",
			},
			want: entity.Property{
				Name:       "getProperty",
				Type:       "Int",
				IsScalar:   true,
				IsNullable: true,
				Parameters: []entity.Property{
					{
						Name:       "name",
						Type:       "String",
						IsScalar:   true,
						IsNullable: true,
					},
				},
			},
		},
		{
			name: "success on property with multiple params",
			args: args{
				input: "getProperty(name: String, age: [Age]!) : Int",
			},
			want: entity.Property{
				Name:       "getProperty",
				Type:       "Int",
				IsScalar:   true,
				IsNullable: true,
				Parameters: []entity.Property{
					{
						Name:       "name",
						Type:       "String",
						IsScalar:   true,
						IsNullable: true,
					},
					{
						Name:   "age",
						Type:   "Age",
						IsList: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			got, err := e.Property(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extractor.Property() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extractor.Property() got \n%+v \nwant \n%+v", got, tt.want)
			}
		})
	}
}

func Test_extractPropTypes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want entity.Property
	}{
		{
			name: "success extract name",
			args: args{
				input: "Image",
			},
			want: entity.Property{
				Name:       "Image",
				IsNullable: true,
			},
		},
		{
			name: "success extract name scalar",
			args: args{
				input: "Int",
			},
			want: entity.Property{
				Name:       "Int",
				IsScalar:   true,
				IsNullable: true,
			},
		},
		{
			name: "success extract name scalar list",
			args: args{
				input: "[Int]",
			},
			want: entity.Property{
				Name:       "Int",
				IsScalar:   true,
				IsNullable: true,
				IsList:     true,
			},
		},
		{
			name: "success extract name scalar list non-nullable",
			args: args{
				input: "[Int]!",
			},
			want: entity.Property{
				Name:       "Int",
				IsScalar:   true,
				IsNullable: false,
				IsList:     true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractPropTypes(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractPropTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
