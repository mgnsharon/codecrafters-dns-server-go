package main

import (
	"bytes"
	"testing"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

func TestHandleRequest(t *testing.T) {
	tcs := []struct {
		n string
		data []byte
		expected  []byte
	}{
		{
			n: "test handleRequest",
			data: req.Bytes(),
			expected: msg.Bytes(),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := handleRequest(tc.data)
			if !bytes.Equal(actual, tc.expected){
				t.Errorf("Expected %v, got %v", string(tc.expected), string(actual))
			}
		})
	}	
}

var msg dns.Message = dns.Message{
	Header: dns.Header{
		ID: 1111,
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
	Questions: []dns.Question{
		{
			Name: dns.DomainName{
				Labels: []dns.DomainLabel{
					{
						Length: 4,
						Content: "test",
					},
					{
						Length: 4,
						Content: "mail",
					},
					{
						Length: 2,
						Content: "io",
					},
				},
			},
			Type: dns.A,
			Class: dns.IN,
		},
						
	},
	Answers: []dns.ResourceRecord{
		{
			Name: dns.DomainName{
				Labels: []dns.DomainLabel{
					{
						Length: 4,
						Content: "test",
					},
					{
						Length: 4,
						Content: "mail",
					},
					{
						Length: 2,
						Content: "io",
					},
				},
			},
			Type: dns.A,
			Class: dns.IN,
			TTL: 60,
			RData: (*dns.IPv4Address)(&dns.IPv4Address{
					Octets: [4]uint8{8, 8, 8, 8},
				}).Bytes(),
			RDLength: 4,
		},
	},
}

var req dns.Message = dns.Message{
	Header: dns.Header{
		ID: 1111,
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
	},
	Questions: []dns.Question{
		{
			Name: dns.DomainName{
				Labels: []dns.DomainLabel{
					{
						Length: 4,
						Content: "test",
					},
					{
						Length: 4,
						Content: "mail",
					},
					{
						Length: 2,
						Content: "io",
					},
				},
			},
			Type: dns.A,
			Class: dns.IN,
		},
						
	},
	
}