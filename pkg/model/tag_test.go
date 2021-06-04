package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	tags := Tags{"a", "b", "c"}

	assert.True(t, tags.Contains("a"))
	assert.False(t, tags.Contains("d"))
}
