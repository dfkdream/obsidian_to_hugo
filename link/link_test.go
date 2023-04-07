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
			name: "inter-directory link",
			fields: fields{
				Reference: "posts2/post5",
				Alt:       "",
				Heading:   "example heading",
				Original:  "",
			},
			want: "[posts2/post5#example heading](posts2/post5#example-heading)",
		},
		{
			name: "unicode link",
			fields: fields{
				Reference: "안녕 hello 세상 world (hi)",
				Alt:       "안녕 world",
				Heading:   "한글 (hangul) 헤딩(heading)  링크 테스트",
				Original:  "",
			},
			want: "[안녕 world](안녕 hello 세상 world (hi)#한글-hangul-헤딩heading--링크-테스트)",
		},
		{
			name: "http link",
			fields: fields{
				Reference: "https://blog.dfkdream.dev",
				Alt:       "",
				Heading:   "(test)",
				Original:  "",
			},
			want: "[https://blog.dfkdream.dev#(test)](https://blog.dfkdream.dev#(test))",
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
