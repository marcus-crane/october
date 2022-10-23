package backend

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitHighlight_Small(t *testing.T) {
	var b strings.Builder
	for i := 0; i < MaxHighlightLen; i++ {
		fmt.Fprintf(&b, "%s", "a")
	}
	expected := []string{b.String()}
	actual := splitHighlight(b.String(), MaxHighlightLen)
	assert.Equal(t, expected, actual)
}

func TestSplitHighlight_OverLimit(t *testing.T) {
	var s string
	for i := 0; i < MaxHighlightLen*3; i++ {
		s += "a"
	}
	firstBit := s[0:MaxHighlightLen]
	secondBit := s[MaxHighlightLen : MaxHighlightLen*2]
	thirdBit := s[MaxHighlightLen*2 : MaxHighlightLen*3]
	expected := []string{firstBit, secondBit, thirdBit}
	actual := splitHighlight(s, MaxHighlightLen)
	assert.Equal(t, expected, actual)
}
