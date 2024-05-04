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

var testQuestion Question = Question{
	Name: DomainName{
		Labels: []DomainLabel{
			{
				Length: 12,
				Content: "codecrafters",
			},
			{
				Length: 2,
				Content: "io",
			},
		},
	},
	Type: A,
	Class: IN,
}

var testResourceRecord ResourceRecord = ResourceRecord{
	Name: DomainName{
		Labels: []DomainLabel{
			{
				Length: 12,
				Content: "codecrafters",
			},
			{
				Length: 2,
				Content: "io",
			},
		},
	},
	Type: A,
	Class: IN,
	TTL: 60,
	RData: (*IPv4Address)(&IPv4Address{
			Octets: [4]uint8{127, 0, 0, 1},
		}).Bytes(),
	RDLength: 4,
}

var testMessage Message = Message{
	Header: testHeader,
	Questions: []Question{
		testQuestion,
	},
	Answers: []ResourceRecord{
		testResourceRecord,
	},
}

var msg Message = Message{
	Header: Header{
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
		ANCount: 1,
		NSCount: 0,
		ARCount: 0,
	},
	Questions: []Question{
		{
			Name: DomainName{
				Labels: []DomainLabel{
					{
						Length: 12,
						Content: "codecrafters",
					},
					{
						Length: 2,
						Content: "io",
					},
				},
			},
			Type: A,
			Class: IN,
		},
						
	},
	Answers: []ResourceRecord{
		{
			Name: DomainName{
				Labels: []DomainLabel{
					{
						Length: 12,
						Content: "codecrafters",
					},
					{
						Length: 2,
						Content: "io",
					},
				},
			},
			Type: A,
			Class: IN,
			TTL: 60,
			RData: (*IPv4Address)(&IPv4Address{
					Octets: [4]uint8{127, 0, 0, 1},
				}).Bytes(),
			RDLength: 4,
		},
	},
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

func TestHeaderFromBytes(t *testing.T) {
	tcs := []struct {
		n string
		header []byte
		expected  uint16
	}{
		{
			n: "test header ID",
			header: testHeader.Bytes(),
			expected: testHeader.ID,
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

func TestMessageFromBytes(t *testing.T) {
	tcs := []struct {
		n string
		message []byte
		expected  uint16
	}{
		{
			n: "test header ID",
			message: msg.Bytes(),
			expected: testMessage.Header.ID,
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