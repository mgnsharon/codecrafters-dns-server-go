package dns

import (
	"testing"
)

func TestMessageFromBytes(t *testing.T) {
	var req Message = loadJson(t, "req", "message").(Message)
	var res Message = loadJson(t, "res", "message").(Message)
	tcs := []struct {
		n string
		message []byte
		expected  uint16
	}{
		{
			n: "test header ID",
			message: req.Bytes(),
			expected: res.Header.ID,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := MessageFromBytes(tc.message)
			if actual.Header.ID != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual.Header.ID)
			}
		})
	}

}