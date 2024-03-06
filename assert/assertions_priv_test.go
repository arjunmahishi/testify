package assert

import (
	"bufio"
	"strings"
	"testing"
)

const (
	// set maxScanTokenSize to 1 more than the default set by bufio. This will
	// cover the case where a single line is longer than the default
	// maxScanTokenSize
	maxScanTokenSize = bufio.MaxScanTokenSize + 1
)

func Test_indentMessageLines(t *testing.T) {
	tt := []struct {
		name            string
		msg             string
		longestLabelLen int
		expected        string
	}{
		{
			name:            "single line",
			msg:             "Hello World\n",
			longestLabelLen: 11,
			expected:        "Hello World",
		},
		{
			name:            "multi line",
			msg:             "Hello\nWorld\n",
			longestLabelLen: 11,
			expected:        "Hello\n\t            \tWorld",
		},
		{
			name:            "single line - extremely long",
			msg:             strings.Repeat("hello ", maxScanTokenSize),
			longestLabelLen: 11,
			expected:        strings.Repeat("hello ", maxScanTokenSize),
		},
		{
			name:            "multi line - extremely long",
			msg:             strings.Repeat("hello\n", maxScanTokenSize),
			longestLabelLen: 3,
			expected: strings.TrimSpace(
				strings.TrimPrefix(strings.Repeat("\thello\n\t    ", maxScanTokenSize), "\t"),
			),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			Equal(t, tc.expected, indentMessageLines(tc.msg, tc.longestLabelLen))
		})
	}
}
