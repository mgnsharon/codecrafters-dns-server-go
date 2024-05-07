package dns

import (
	"testing"
)

func TestMessageFromBytes(t *testing.T) {
	
	tcs := []struct {
		n string
		message []byte
		expected  uint16
	}{
		{
			n: "fwd msg",
			// "\x8eÁ\x80\x00\x01\x00\x01\x00\x00\x00\x00\fcodecrafters\x02io\x00\x00\x01\x00\x01\xc0\f\x00\x01\x00\x01\x00\x00\x05t\x00\x04LL\x15\x15"
			message: []byte{142,195,129,128,0,1,0,1,0,0,0,0,12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0,0,1,0,1,192,12,0,1,0,1,0,0,5,116,0,4,76,76,21,21},
			expected: 1234,
		},
		{
			n: "uncompressed msg",
			// "K%\x01\x00\x00\x02\x00\x00\x00\x00\x00\x00\x03abc\x11longassdomainname\x03com\x00\x00\x01\x00\x01\x03def\xc0\x10\x00\x01\x00\x01"

			message: []byte{75,37,1,0,0,2,0,0,0,0,0,0,3,97,98,99,17,108,111,110,103,97,115,115,100,111,109,97,105,110,110,97,109,101,3,99,111,109,0,0,1,0,1,3,100,101,102,192,16,0,1,0,1},
			expected: 9533,
		},
		
		
		/* {
			n: "test header ID",
			message: req.Bytes(),
			expected: res.Header.ID,
		},
		{
			n: "uncompressed msg",
			// "F\xbd\x01\x00\x00\x01\x00\x00\x00\x00\x00\x00\fcodecrafters\x02io\x00\x00\x01\x00\x01"
			message: []byte{70, 189, 1 ,0,0,1,0,0,0,0,0,0,12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0,0,1,0,1},
			expected: 1234,
		}, */
		
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