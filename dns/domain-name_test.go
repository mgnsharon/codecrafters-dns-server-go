package dns

import (
	"testing"
)

func TestDomainLabelFromBytes(t *testing.T) {
	tcs := []struct {
		n string
		data []byte
		expected DomainLabel
	}{
		{
			n: "test DomainLabelFromBytes",
			data: []byte{12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111},
			expected: DomainLabel{Length: 12, Content: "codecrafters"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := DomainLabelFromBytes(tc.data)
			if actual.Content != tc.expected.Content {
				t.Errorf("Expected %v, got %v", tc.expected.Content, actual.Content)
			}
		})
	}
}

func TestDomainNameFromBytes(t *testing.T) {
	tcs := []struct {
		n string
		data []byte
		msgData []byte
		expected DomainName
	}{
		{
			n: "DomainNameFromBytes: codecrafters.io",
			data: []byte{12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0},
			msgData: []byte{142,195,1,0,0,1,0,0,0,0,0,0,12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0,0,1,0,1},
			expected: DomainName{Labels: []DomainLabel{{Length: 12, Content: "codecrafters"}, {Length: 2, Content: "io"}}},
		},
		{
			n: "DomainNameFromBytes: abc.longassdomainname.com",
			data: []byte{3,97,98,99,17,108,111,110,103,97,115,115,100,111,109,97,105,110,110,97,109,101,3,99,111,109,0,0,1,0,1,3,100,101,102,192,16,0,1,0,1},
			msgData: []byte{75,37,1,0,0,2,0,0,0,0,0,0,3,97,98,99,17,108,111,110,103,97,115,115,100,111,109,97,105,110,110,97,109,101,3,99,111,109,0,0,1,0,1,3,100,101,102,192,16,0,1,0,1},
			expected: DomainName{Labels: []DomainLabel{{Length: 3, Content: "abc"}, {Length: 17, Content: "longassdomainname"}, {Length: 3, Content: "com"}}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := DomainNameFromBytes(tc.data, tc.msgData)
			if actual.Labels[0].Content != tc.expected.Labels[0].Content {
				t.Errorf("Expected %v, got %v", tc.expected.Labels[0].Content, actual.Labels[0].Content)
			}
		})
	}
}

func TestDomainNameBytes(t *testing.T) {
	tcs := []struct {
		n string
		data DomainName
		expected []byte
	}{
		{
			n: "DomainNameBytes: codecrafters.io",
			data: DomainName{Labels: []DomainLabel{{Length: 12, Content: "codecrafters"}, {Length: 2, Content: "io"}}},
			expected: []byte{12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0},
		},
		{
			n: "DomainNameBytes: abc.longassdomainname.com",
			data: DomainName{Labels: []DomainLabel{{Length: 3, Content: "abc"}, {Length: 17, Content: "longassdomainname"}, {Length: 3, Content: "com"}}},
			expected: []byte{3,97,98,99,17,108,111,110,103,97,115,115,100,111,109,97,105,110,110,97,109,101,3,99,111,109,0},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := tc.data.Bytes()
			if string(actual) != string(tc.expected) {
				t.Errorf("Expected %v, got %v", string(tc.expected), string(actual))
			}
		})
	}
}