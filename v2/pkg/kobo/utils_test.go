package kobo

import (
	"reflect"
	"testing"
)

func TestGetRelativeKoboPath(t *testing.T) {
	tests := []struct {
		expected string
		input    string
	}{
		{
			expected: "Fadell, Tony/Build - Tony Fadell.kepub.epub",
			input:    "file:///mnt/onboard/Fadell, Tony/Build - Tony Fadell.kepub.epub",
		},
		{
			expected: "Monteiro, Mike/Ruined by Design_ How Designers Destroyed the World, and What We Can Do to Fix It - Mike Monteiro.kepub.epub",
			input:    "file:///mnt/onboard/Monteiro, Mike/Ruined by Design_ How Designers Destroyed the World, and What We Can Do to Fix It - Mike Monteiro.kepub.epub",
		},
		{
			expected: "Grove, Andrew S_/High Output Management - Andrew S. Grove.kepub.epub",
			input:    "file:///mnt/onboard/Grove, Andrew S_/High Output Management - Andrew S. Grove.kepub.epub",
		},
		{
			expected: "Vend/Technology at Vend - Vend.epub#(2)OEBPS/_projects_work.xhtml",
			input:    "file:///mnt/onboard/Vend/Technology at Vend - Vend.epub#(2)OEBPS/_projects_work.xhtml",
		},
		{
			expected: "Herron, Mick/Slow Horses - Mick Herron.epub#(8)OEBPS/9781569476437_Split12.html",
			input:    "file:///mnt/onboard/Herron, Mick/Slow Horses - Mick Herron.epub#(8)OEBPS/9781569476437_Split12.html",
		},
		{
			expected: "Fadell, Tony/Build - Tony Fadell.kepub.epub!!OEBPS/text/9780063046078_Chapter_16.xhtml",
			input:    "/mnt/onboard/Fadell, Tony/Build - Tony Fadell.kepub.epub!!OEBPS/text/9780063046078_Chapter_16.xhtml",
		},
	}

	for i, tc := range tests {
		actual := getRelativeKoboPath(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.expected, actual)
		}
	}
}

func TestTrimContentFileName(t *testing.T) {
	tests := []struct {
		expected string
		input    string
	}{
		{
			expected: "OEBPS/text/9780063046078_Chapter_16.xhtml",
			input:    "Fadell, Tony/Build - Tony Fadell.kepub.epub!!OEBPS/text/9780063046078_Chapter_16.xhtml",
		},
		{
			expected: "(2)OEBPS/_projects_work.xhtml",
			input:    "Vend/Technology at Vend - Vend.epub#(2)OEBPS/_projects_work.xhtml",
		},
		{
			expected: "(8)OEBPS/9781569476437_Split12.html",
			input:    "file:///mnt/onboard/Herron, Mick/Slow Horses - Mick Herron.epub#(8)OEBPS/9781569476437_Split12.html",
		},
		{
			expected: "epub/text/chapter-1.xhtml",
			input:    "Marx, Karl & Engles, Fridrick/Communist Manifesto, The - Karl Marx & Fridrick Engles.kepub.epub!!epub/text/chapter-1.xhtml",
		},
	}

	for i, tc := range tests {
		actual := trimContentFileName(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.expected, actual)
		}
	}
}
