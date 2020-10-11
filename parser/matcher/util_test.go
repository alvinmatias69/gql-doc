package matcher

import "testing"

func Test_formatNamed(t *testing.T) {
	type args struct {
		format string
		entry  map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success generate named string",
			args: args{
				format: "hello {{.name}}",
				entry: map[string]string{
					"name": "there",
				},
			},
			want: "hello there",
		},
		{
			name: "failure on bad template",
			args: args{
				format: "hello {{.name}",
				entry: map[string]string{
					"name": "there",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatNamed(tt.args.format, tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatNamed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("formatNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}
