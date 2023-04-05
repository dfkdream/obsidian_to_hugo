package link

import "strings"

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
	result += escapeString(l.Reference)
	if l.Heading != "" {
		result += "#" + l.Heading
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

func isMarkdownEligibleCharacter(c rune) bool {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

	if strings.ContainsRune(chars, c) {
		return true
	}

	return false
}

func escapeString(s string) string {
	result := ""
	needToSkip := false
	for _, c := range s {
		if !isMarkdownEligibleCharacter(c) {
			if !needToSkip {
				needToSkip = true
				result += "-"
			}
			continue
		}

		result += string(c)
		needToSkip = false
	}

	return strings.Trim(result, "-")
}
