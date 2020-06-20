package extractor

// func TestComment(t *testing.T) {
// 	type args struct {
// 		input string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "success extract comment",
// 			args: args{
// 				input: "# this is a comment",
// 			},
// 			want: "this is a comment\n",
// 		},
// 		{
// 			name: "success extract comment inline",
// 			args: args{
// 				input: "age: Int # this is a comment",
// 			},
// 			want: "this is a comment\n",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Comment(tt.args.input); got != tt.want {
// 				t.Errorf("Comment() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
