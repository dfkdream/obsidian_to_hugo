package link

import (
	"reflect"
	"testing"
)

func TestMarkdownLinkFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []Link
	}{
		{
			name: "reference with alt",
			args: args{
				s: "[hello world](hello world)",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello world",
					Heading:   "",
				},
			},
		},
		{
			name: "reference with heading and alt",
			args: args{
				s: "[hello world hi](hello world#hi)",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello world hi",
					Heading:   "hi",
				},
			},
		},
		{
			name: "external link",
			args: args{
				s: "[external](https://blog.dfkdream.dev/)",
			},
			want: []Link{
				{
					Reference: "https://blog.dfkdream.dev/",
					Alt:       "external",
					Heading:   "",
				},
			},
		},
		{
			name: "multiple markdown links",
			args: args{
				s: "[hello world](hello world)[X][ ][[]][external](https://blog.dfkdream.dev/)hi[hello world hi](hello world#hi)",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello world",
					Heading:   "",
				},
				{
					Reference: "https://blog.dfkdream.dev/",
					Alt:       "external",
					Heading:   "",
				},
				{
					Reference: "hello world",
					Alt:       "hello world hi",
					Heading:   "hi",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarkdownLinkFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownLinkFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
