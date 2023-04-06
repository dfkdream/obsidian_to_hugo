package link

type Link struct {
	Reference string
	Alt       string
	Heading   string
	Original  string
}

func (l Link) MarkdownLink() string {
	result := "["
	if l.Alt == "" {
		l.Alt = l.Reference
		if l.Heading != "" {
			l.Alt += "#" + l.Heading
		}
	}
	result += l.Alt + "]("
	result += sanitizeString(l.Reference)
	if l.Heading != "" {
		result += "#" + sanitizeString(l.Heading)
	}
	result += ")"
	return result
}

func (l Link) WikiLink() string {
	result := "[[" + l.Reference
	if l.Heading != "" {
		result += "#" + l.Heading
	}
	if l.Alt != "" {
		result += "|" + l.Alt
	}
	result += "]]"
	return result
}
