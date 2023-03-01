package link

import (
	"reflect"
	"testing"
)

func TestWikiLinkFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []Link
	}{
		{
			name: "reference only",
			args: args{
				s: "[[hello world]]",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "",
					Heading:   "",
				},
			},
		},
		{
			name: "reference with alt",
			args: args{
				s: "[[hello world|hello]]",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello",
					Heading:   "",
				},
			},
		},
		{
			name: "reference with heading",
			args: args{
				s: "[[hello world#hi]]",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "",
					Heading:   "hi",
				},
			},
		},
		{
			name: "reference with heading and alt",
			args: args{
				s: "[[hello world#hi|hello]]",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello",
					Heading:   "hi",
				},
			},
		},
		{
			name: "multiple wikilinks",
			args: args{
				s: "[[hello world#hi|hello]] hello world[][[][]][[foo|bar]][[lorem#ipsum|dolor]]",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello",
					Heading:   "hi",
				},
				{
					Reference: "foo",
					Alt:       "bar",
					Heading:   "",
				},
				{
					Reference: "lorem",
					Alt:       "dolor",
					Heading:   "ipsum",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WikiLinkFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WikiLinkFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
