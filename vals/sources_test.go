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
			name:     "hard-coded_JOF",
			val:      "JOF",
			expected: true,
		},
		{
			name:     "hard-coded_ING",
			val:      "ING",
			expected: true,
		},
		{
			name:     "constant_JOF",
			val:      ScraperSourceJofogasHu,
			expected: true,
		},
		{
			name:     "constant_ING",
			val:      ScraperSourceIngatlanCom,
			expected: true,
		},
		{
			name:     "lowercase_ing",
			val:      "ing",
			expected: false,
		},
		{
			name:     "lowercase_jof",
			val:      "jof",
			expected: false,
		},
		{
			name:     "random_1",
			val:      "asd",
			expected: false,
		},
		{
			name:     "random_2",
			val:      "ASD",
			expected: false,
		},
		{
			name:     "empty_string",
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
