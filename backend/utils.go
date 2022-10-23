package backend

import "strings"

func splitHighlight(highlight string, chunkSize int) []string {
	if chunkSize >= len(highlight) {
		return []string{highlight}
	}
	var chunks []string
	var b strings.Builder
	b.Grow(chunkSize)
	l := 0
	for _, r := range highlight {
		b.WriteRune(r)
		l++
		if l == chunkSize {
			chunks = append(chunks, b.String())
			l = 0
			b.Reset()
			b.Grow(chunkSize)
		}
	}
	if l > 0 {
		chunks = append(chunks, b.String())
	}
	return chunks
}
