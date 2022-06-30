package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertString2Slice(t *testing.T) {
	in := "a,b,c,d"
	res := ConvertString2slice(in, ",")
	assert.NotNil(t, res)
	assert.Equal(t, []string{"a", "b", "c", "d"}, res)
}
