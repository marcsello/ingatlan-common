package vals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsScraperSourceValid(t *testing.T) {

	type testCase struct {
		name     string
		val      string
		expected bool
	}

	testCases := []testCase{
		{
			name:     "hard-coded JOF",
			val:      "JOF",
			expected: true,
		},
		{
			name:     "hard-coded ING",
			val:      "ING",
			expected: true,
		},
		{
			name:     "constant JOF",
			val:      ScraperSourceJofogasHu,
			expected: true,
		},
		{
			name:     "constant ING",
			val:      ScraperSourceIngatlanCom,
			expected: true,
		},
		{
			name:     "lowercase ing",
			val:      "ing",
			expected: false,
		},
		{
			name:     "lowercase jof",
			val:      "jof",
			expected: false,
		},
		{
			name:     "random 1",
			val:      "asd",
			expected: false,
		},
		{
			name:     "random 2",
			val:      "ASD",
			expected: false,
		},
		{
			name:     "empty string",
			val:      "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.expected {
				assert.True(t, IsScraperSourceValid(tc.val))
			} else {
				assert.False(t, IsScraperSourceValid(tc.val))
			}

		})
	}

}
