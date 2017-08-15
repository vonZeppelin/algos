package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	maxLen, curLen, seen, max := 0, 0, make(map[rune]int), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, chr := range s {
		if j, ok := seen[chr]; ok {
			maxLen = max(maxLen, curLen)
			curLen = i - j
		} else {
			curLen++
		}
		seen[chr] = i
	}

	return max(maxLen, curLen)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, lengthOfLongestSubstring("c"))
	assert.Equal(1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(3, lengthOfLongestSubstring("dvdf"))
	assert.Equal(2, lengthOfLongestSubstring("abba"))
	assert.Equal(0, lengthOfLongestSubstring(""))
}
