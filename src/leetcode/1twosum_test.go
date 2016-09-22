package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	idxs := make(map[int]int)
	for i, num := range nums {
		if j, ok := idxs[target-num]; ok {
			return []int{j, i}
		}
		idxs[num] = i
	}
	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal([]int{2, 3}, twoSum([]int{2, 7, 11, 15}, 26))
	assert.Equal([]int{-1, -1}, twoSum([]int{2, 7, 11, 15}, 1))
	assert.Equal([]int{-1, -1}, twoSum([]int{}, 1))
}
