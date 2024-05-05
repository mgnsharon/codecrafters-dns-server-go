package dns

import (
	"testing"
)

func TestHeader(t *testing.T) {
	var th Header = loadJson(t, "header", "header").(Header)
	tcs := []struct {
		header Header
		expected  int
	}{
		{
			header: th,
			expected: 12,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			if len(tc.header.Bytes()) != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, len(tc.header.Bytes()))
			}
		})
	}

}

func TestHeaderFromBytes(t *testing.T) {
	var th Header = loadJson(t, "header", "header").(Header)
	tcs := []struct {
		n string
		header []byte
		expected  uint16
	}{
		{
			n: "test header ID",
			header: th.Bytes(),
			expected: th.ID,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := HeaderFromBytes(tc.header)
			if actual.ID != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual.ID)
			}
		})
	}

}
