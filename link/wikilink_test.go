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
					Original:  "[[hello world]]",
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
					Original:  "[[hello world|hello]]",
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
					Original:  "[[hello world#hi]]",
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
					Original:  "[[hello world#hi|hello]]",
				},
			},
		},
		{
			name: "multiple wikilinks",
			args: args{
				s: "[[hello world#hi|hello]] hello world[][[][]][[foo|bar]][X][ ][[lorem#ipsum|dolor]]",
			},
			want: []Link{
				{
					Reference: "hello world",
					Alt:       "hello",
					Heading:   "hi",
					Original:  "[[hello world#hi|hello]]",
				},
				{
					Reference: "foo",
					Alt:       "bar",
					Heading:   "",
					Original:  "[[foo|bar]]",
				},
				{
					Reference: "lorem",
					Alt:       "dolor",
					Heading:   "ipsum",
					Original:  "[[lorem#ipsum|dolor]]",
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
