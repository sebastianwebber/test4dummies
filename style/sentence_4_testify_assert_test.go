package style

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testify_assert_Sentence(t *testing.T) {

	assert := assert.New(t)

	out, err := Sentence(input)

	assert.Nilf(err, "should work without errors")
	assert.Equal(out, expected, "should convert strings properly")
}
