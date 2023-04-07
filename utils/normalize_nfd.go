package utils

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func NormalizeNFD(s string) string {
	t := transform.Chain(norm.NFD, norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}
