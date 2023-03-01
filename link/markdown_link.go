package link

import (
	"regexp"
	"strings"
)

func MarkdownLinkFromString(s string) []Link {
	// TODO: Need to find something better than regex
	re := regexp.MustCompile(`\[([^\[^\]]*?)]\((.*?)\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	result := make([]Link, len(matches))

	for i, v := range matches {
		result[i].Alt = v[1]

		headingIndex := strings.LastIndex(v[2], "#")
		ref := v[2]
		if headingIndex > -1 {
			result[i].Heading = ref[headingIndex+1:]
			ref = ref[:headingIndex]
		}

		result[i].Reference = ref
	}

	return result
}
