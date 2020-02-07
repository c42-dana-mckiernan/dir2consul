package files

import (
	"testing"
)

func TestIsFileWanted(t *testing.T) {
	fixtures := []struct {
		filename   string
		extensions []string
		result     bool
	}{
		{"foo.bar", []string{".txt", ".bar"}, true},
		{"foo.txt", []string{".txt", ".bar"}, true},
		{".bar", []string{".txt", ".bar"}, true},
		{"foo.json", []string{".txt", ".bar"}, false},
		{"foo.pdf", []string{".txt", ".bar"}, false},
		{"bar.xyz.xyz", []string{".txt", ".bar"}, false},
		{"foo", []string{".txt", ".bar"}, false},
	}

	for _, s := range fixtures {
		if r := isFileWanted(s.filename, s.extensions); r != s.result {
			t.Error("Failed to match")
		}
	}
}
