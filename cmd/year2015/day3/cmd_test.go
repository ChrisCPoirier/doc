package day3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	b, err := os.ReadFile(`test.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected int
		input    []byte
		fn       func([]byte) int
	}{
		{
			expected: 4,
			input:    b,
			fn:       part1,
		},
		{
			expected: 12,
			input:    b,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}
