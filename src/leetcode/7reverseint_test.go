package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func helper(x, acc int) int {
	if x < 10 {
		return x * acc
	} else {
		return x%10*acc + helper(x/10, acc*10)
	}
}

func reverse(x int) int {
	return helper(x, 1)
}

func TestReverse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(321, reverse(123))
	assert.Equal(-321, reverse(-123))
}
