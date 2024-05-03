package dns

import (
	"testing"
)

var testHeader Header =  Header{
	ID: 1234,
	QR: 1,
	Opcode: 0,
	AA: 0,
	TC: 0,
	RD: 0,
	RA: 0,
	Z: 0,
	RCode: 0,
	QDCount: 1,
	ANCount: 0,
	NSCount: 0,
	ARCount: 0,
}

func TestHeader(t *testing.T) {
	tcs := []struct {
		header Header
		expected  int
	}{
		{
			header: testHeader,
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