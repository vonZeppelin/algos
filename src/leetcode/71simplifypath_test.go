package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func simplifyPath(path string) string {
	var segments []string
	var segment []rune

	for i, runes, size := 0, []rune(path), len(path); i < size; {
		switch runes[i] {
		case '/':
			if len(segment) > 0 {
				segments = append(segments, string(segment))
				segment = nil
			}
			i++
		case '.':
			dotCount := 1
			for i++; i < size && runes[i] == '.'; i++ {
				dotCount++
			}
			if dotCount == 2 {
				lastSegmentIdx := len(segments) - 1
				if lastSegmentIdx > -1 {
					segments = segments[:lastSegmentIdx]
				}
			} else if dotCount > 2 {
				segments = append(segments, string(runes[i-dotCount:i]))
			}
		default:
			segment = append(segment, runes[i])
			i++
		}
	}

	return "/" + strings.Join(segments, "/")
}

func TestSimplifyPath(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("/home", simplifyPath("/home/"))
	assert.Equal("/c", simplifyPath("/a/./b/../../c/"))
	assert.Equal("/", simplifyPath("/../"))
	assert.Equal("/home/foo", simplifyPath("/home//foo/"))
	assert.Equal("/...", simplifyPath("/..."))
	assert.Equal("/..hidden", simplifyPath("/..hidden"))
}
