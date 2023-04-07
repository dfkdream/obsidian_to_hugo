package link

import "strings"

type Link struct {
	Reference string
	Alt       string
	Heading   string
	Original  string
}

func (l Link) isInternal() bool {
	if strings.HasPrefix(l.Reference, "http://") ||
		strings.HasPrefix(l.Reference, "https://") {
		return false
	}
	return true
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
	result += l.Reference
	if l.Heading != "" {
		if l.isInternal() {
			result += "#" + sanitizeString(l.Heading)
		} else {
			result += "#" + l.Heading
		}
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
