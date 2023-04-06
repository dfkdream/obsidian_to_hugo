// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Reference: https://github.com/gohugoio/hugo/blob/f1e8f010f5f5990c6e172b977e5e2d2878b9a338/markup/goldmark/autoid.go

package link

import (
	"strings"
	"unicode"
)

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func sanitizeString(s string) string {
	var result strings.Builder
	for _, r := range s {
		switch {
		case r == '-' || r == ' ':
			result.WriteRune('-')
		case isAlphaNumeric(r):
			result.WriteRune(r)
		default:
		}
	}

	return result.String()
}
