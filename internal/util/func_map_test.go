package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPtr(t *testing.T) {
	// Sample strings
	strings := []string{
		"*This is a test string.",
		"This string does not start with an asterisk.",
		"*Another example with an asterisk.",
		"No asterisk here.",
	}

	// Matching strings
	var matchedStrings []string
	for _, s := range strings {
		if IsPtr(s) {
			matchedStrings = append(matchedStrings, s)
		}
	}

	require.Equal(t, matchedStrings, []string{"*This is a test string.",
		"*Another example with an asterisk."})
}
