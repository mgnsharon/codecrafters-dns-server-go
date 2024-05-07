package dns

import (
	"bytes"
	"testing"
)

func TestResourceRecordFromBytes(t *testing.T) {
	tcs := []struct {
		n string
		data []byte
		offset int
		expected ResourceRecord
	}{
		{
			n: "ResourceRecordFromBytes: uncompressed codecrafters.io",
			data: []byte{142,195,1,0,0,1,0,0,0,0,0,0,12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0,0,1,0,1,192,12,0,1,0,4,76,76,21,21},
			offset: 12,
			expected: ResourceRecord{
				Name: DomainName{Labels: []DomainLabel{{Length: 12, Content: "codecrafters"}, {Length: 2, Content: "io"}}},
				Type: RecordType(1),
				Class: RecordClass(1),
				TTL: 3222011905,
				RDLength: 4,
				RData: []byte{76, 76, 21, 21},
			},
		},
		{
			n: "ResourceRecordFromBytes: uncompressed abc.longassdomainname.com",
			data: []byte{75,37,1,0,0,2,0,0,0,0,0,0,3,97,98,99,17,108,111,110,103,97,115,115,100,111,109,97,105,110,110,97,109,101,3,99,111,109,0,0,1,0,1,3,100,101,102,0,4,192,16,0,1,},
			offset: 12,
			expected: ResourceRecord{
				Name: DomainName{Labels: []DomainLabel{{Length: 3, Content: "abc"}, {Length: 17, Content: "longassdomainname"}, {Length: 3, Content: "com"}}},
				Type: RecordType(1),
				Class: RecordClass(1),
				TTL: 56911206,
				RDLength: 4,
				RData: []byte{192, 16, 0, 1},
			},
		},
		{
			n: "ResourceRecordFromBytes: compressed def.longassdomain.com",
			data: []byte{75,37,1,0,0,2,0,0,0,0,0,0,3,97,98,99,17,108,111,110,103,97,115,115,100,111,109,97,105,110,110,97,109,101,3,99,111,109,0 ,0,1,0,1,3,100,101,102,192,16,0,0,1,0,1,3,100,101,102,0,4,192,16,0,1,},
			offset: 43,
			expected: ResourceRecord{
				Name: DomainName{Labels: []DomainLabel{{Length: 3, Content: "def"}, {Length: 17, Content: "longassdomainname"}, {Length: 3, Content: "com"}}},
				Type: RecordType(1),
				Class: RecordClass(1),
				TTL: 56911206,
				RDLength: 4,
				RData: []byte{192, 16, 0, 1},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := ResourceRecordFromBytes(tc.data[tc.offset:], tc.data)
			if actual.Name.Labels[0].Content != tc.expected.Name.Labels[0].Content {
				t.Errorf("Expected %v, got %v", tc.expected.Name.Labels[0].Content, actual.Name.Labels[0].Content)
			}
			if actual.Class != tc.expected.Class {
				t.Errorf("Expected %v, got %v", tc.expected.Class, actual.Class)
			}
			if actual.Type != tc.expected.Type {
				t.Errorf("Expected %v, got %v", tc.expected.Type, actual.Type)
			}
			if actual.TTL != tc.expected.TTL {
				t.Errorf("Expected %v, got %v", tc.expected.TTL, actual.TTL)
			}
			if actual.RDLength != tc.expected.RDLength {
				t.Errorf("Expected %v, got %v", tc.expected.RDLength, actual.RDLength)
			}
			if bytes.Equal(actual.RData, tc.expected.RData) {
				t.Errorf("Expected %v, got %v", tc.expected.RData, actual.RData)
			}
		})
	}
}