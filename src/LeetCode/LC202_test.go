package LeetCode_test

import (
	"LeetCode"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestIsHappy(t *testing.T)  {
	assert.Equal(t, true, LeetCode.IsHappy(19))
}
