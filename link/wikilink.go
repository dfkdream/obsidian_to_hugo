package link

import (
	"regexp"
	"strings"
)

func WikiLinkFromString(s string) []Link {
	re := regexp.MustCompile(`\[\[([^]]+)]]`)
	matches := re.FindAllStringSubmatch(s, -1)
	result := make([]Link, len(matches))

	for i, v := range matches {
		altIndex := strings.LastIndex(v[1], "|")
		ref := v[1]
		if altIndex > -1 {
			result[i].Alt = v[1][altIndex+1:]
			ref = v[1][:altIndex]
		}

		headingIndex := strings.LastIndex(ref, "#")
		if headingIndex > -1 {
			result[i].Heading = ref[headingIndex+1:]
			ref = ref[:headingIndex]
		}

		result[i].Reference = ref
	}

	return result
}
