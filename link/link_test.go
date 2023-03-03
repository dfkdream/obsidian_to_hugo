package link

import "testing"

func TestLink_MarkdownLink(t *testing.T) {
	type fields struct {
		Reference string
		Alt       string
		Heading   string
		Original  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "reference only",
			fields: fields{
				Reference: "hello",
				Alt:       "",
				Heading:   "",
				Original:  "",
			},
			want: "[hello](hello)",
		},
		{
			name: "reference only with heading",
			fields: fields{
				Reference: "hello",
				Alt:       "",
				Heading:   "hi",
				Original:  "",
			},
			want: "[hello#hi](hello#hi)",
		},
		{
			name: "reference with alt",
			fields: fields{
				Reference: "hello",
				Alt:       "world",
				Heading:   "",
				Original:  "",
			},
			want: "[world](hello)",
		},
		{
			name: "reference, alt and heading",
			fields: fields{
				Reference: "hello",
				Alt:       "world",
				Heading:   "hi",
				Original:  "",
			},
			want: "[world](hello#hi)",
		},
		{
			name: "markdown-invalid link",
			fields: fields{
				Reference: "hello world (hi)",
				Alt:       "world",
				Heading:   "hi",
				Original:  "",
			},
			want: "[world](hello-world-hi#hi)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{
				Reference: tt.fields.Reference,
				Alt:       tt.fields.Alt,
				Heading:   tt.fields.Heading,
				Original:  tt.fields.Original,
			}
			if got := l.MarkdownLink(); got != tt.want {
				t.Errorf("MarkdownLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_WikiLink(t *testing.T) {
	type fields struct {
		Reference string
		Alt       string
		Heading   string
		Original  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "reference only",
			fields: fields{
				Reference: "hello",
				Alt:       "",
				Heading:   "",
				Original:  "",
			},
			want: "[[hello]]",
		},
		{
			name: "reference only with heading",
			fields: fields{
				Reference: "hello",
				Alt:       "hello",
				Heading:   "hi",
				Original:  "",
			},
			want: "[[hello#hi|hello]]",
		},
		{
			name: "reference with alt",
			fields: fields{
				Reference: "hello",
				Alt:       "world",
				Heading:   "",
				Original:  "",
			},
			want: "[[hello|world]]",
		},
		{
			name: "reference, alt and heading",
			fields: fields{
				Reference: "hello",
				Alt:       "world",
				Heading:   "hi",
				Original:  "",
			},
			want: "[[hello#hi|world]]",
		},
		{
			name: "markdown-invalid link",
			fields: fields{
				Reference: "hello world (hi)",
				Alt:       "world",
				Heading:   "hi",
				Original:  "",
			},
			want: "[[hello world (hi)#hi|world]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{
				Reference: tt.fields.Reference,
				Alt:       tt.fields.Alt,
				Heading:   tt.fields.Heading,
				Original:  tt.fields.Original,
			}
			if got := l.WikiLink(); got != tt.want {
				t.Errorf("WikiLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
