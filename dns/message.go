package dns

import (
	"encoding/binary"
	"fmt"
)



type RecordType uint16 
type RecordClass uint16

const (
	A     RecordType = 1 // A is a host address.
	NS    RecordType = 2 // NS is an authoritative name server.
	CNAME RecordType = 5 // CNAME is the canonical name for an alias.
	SOA   RecordType = 6 // SOA is the start of a zone of authority.
	PTR   RecordType = 12 // PTR is a domain name pointer.
	MX    RecordType = 15 // MX is a mail exchange.
	TXT   RecordType = 16 // TXT is text strings.
	AAAA  RecordType = 28 // AAAA is a host address.

	IN RecordClass = 1 // IN is the Internet class.
	CH RecordClass = 3 // CH is the Chaos class.
	HS RecordClass = 4 // HS is the Hesiod class.
)

// Header represents the DNS message header.
type Header struct {
	ID      uint16 // ID is the identification number of the DNS message.
	QR      uint8  // QR is the query/response flag.
	Opcode  uint8  // Opcode is the operation code.
	AA      uint8  // AA is the authoritative answer flag.
	TC      uint8  // TC is the truncation flag.
	RD      uint8  // RD is the recursion desired flag.
	RA      uint8  // RA is the recursion available flag.
	Z       uint8  // Z is reserved for future use.
	RCode   uint8  // RCode is the response code.
	QDCount uint16 // QDCount is the number of questions in the question section.
	ANCount uint16 // ANCount is the number of resource records in the answer section.
	NSCount uint16 // NSCount is the number of name server resource records in the authority section.
	ARCount uint16 // ARCount is the number of resource records in the additional section.
}

type DomainLabel struct {
	Length uint8
	Content  string
}

type DomainName struct {
	Labels []DomainLabel
}

type Question struct {
	Name DomainName
	Type RecordType
	Class RecordClass
}

type ResourceRecord struct {
	Name DomainName
	Type RecordType
	Class RecordClass
	TTL uint32
	RDLength uint16
	RData []byte
}

type Authority struct {
}

type Additional struct {
}

type Message struct {
	Header     Header
	Questions  []Question
	Answers    []ResourceRecord
	Authorities []Authority
	Additionals []Additional
}

type IPv4Address struct {
	Octets [4]uint8
}

func (a *IPv4Address) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.Octets[0], a.Octets[1], a.Octets[2], a.Octets[3])
}

func (a *IPv4Address) Bytes() []byte {
	buf := make([]byte, 4)
	copy(buf, a.Octets[:])
	return buf
}

// Bytes returns the 12 byte representation of the DNS header.
// BigEndian is used for encoding.
func (h *Header) Bytes() []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:2], h.ID)
	flag := uint16(h.QR)<<15 | uint16(h.Opcode)<<11 | uint16(h.AA)<<10 | uint16(h.TC)<<9 | uint16(h.RD)<<8 | uint16(h.RA)<<7 | uint16(h.Z)<<4 | uint16(h.RCode)
	binary.BigEndian.PutUint16(buf[2:4], flag)
	binary.BigEndian.PutUint16(buf[4:6], h.QDCount)
	binary.BigEndian.PutUint16(buf[6:8], h.ANCount)
	binary.BigEndian.PutUint16(buf[8:10], h.NSCount)
	binary.BigEndian.PutUint16(buf[10:12], h.ARCount)

	return buf
}

func (m *Message) Bytes() []byte {
	buf := m.Header.Bytes()
	for _, question := range m.Questions {
		buf = append(buf, question.Bytes()...)
	}
	for _, answer := range m.Answers {
		buf = append(buf, answer.Bytes()...)
	}
	return buf
}

func (d *DomainLabel) Bytes() []byte {
	buf := make([]byte, 1 + len(d.Content))
	buf[0] = d.Length
	copy(buf[1:], []byte(d.Content))
	return buf
}

func (d *DomainName) Bytes() []byte {
	buf := make([]byte, 0)
	for _, label := range d.Labels {
		buf = append(buf, label.Bytes()...)
	}
	buf = append(buf, byte(0))
	return buf
}

func (q *Question) Bytes() []byte {
	buf := make([]byte, 0)
	buf = append(buf, q.Name.Bytes()...)
	buf = binary.BigEndian.AppendUint16(buf, uint16(q.Type))
	buf = binary.BigEndian.AppendUint16(buf, uint16(q.Class))
	return buf
}

func (r *ResourceRecord) Bytes() []byte {
	buf := make([]byte, 0)
	buf = append(buf, r.Name.Bytes()...)
	buf = binary.BigEndian.AppendUint16(buf, uint16(r.Type))
	buf = binary.BigEndian.AppendUint16(buf, uint16(r.Class))
	buf = binary.BigEndian.AppendUint32(buf, r.TTL)
	buf = binary.BigEndian.AppendUint16(buf, r.RDLength)
	buf = append(buf, r.RData...)
	return buf
}

func HeaderFromBytes(buf []byte) *Header {
	h := Header{}
	h.ID = binary.BigEndian.Uint16(buf[0:2])
	flag := binary.BigEndian.Uint16(buf[2:4])
	h.QR = uint8(flag >> 15)
	h.Opcode = uint8((flag >> 11) & 0x0F)
	h.AA = uint8((flag >> 10) & 0x01)
	h.TC = uint8((flag >> 9) & 0x01)
	h.RD = uint8((flag >> 8) & 0x01)
	h.RA = uint8((flag >> 7) & 0x01)
	h.Z = uint8((flag >> 4) & 0x07)
	h.RCode = uint8(flag & 0x0F)
	h.QDCount = binary.BigEndian.Uint16(buf[4:6])
	h.ANCount = binary.BigEndian.Uint16(buf[6:8])
	h.NSCount = binary.BigEndian.Uint16(buf[8:10])
	h.ARCount = binary.BigEndian.Uint16(buf[10:12])
	return &h
}



func DomainLabelFromBytes(buf []byte) *DomainLabel {
	d := DomainLabel{}
	d.Length = buf[0]
	d.Content = string(buf[1:d.Length + 1])
	return &d
}

func DomainNameFromBytes(buf []byte) *DomainName {
	d := DomainName{}
	d.Labels = make([]DomainLabel, 0)
	for {
		if buf[0] == 0 {
			break
		}
		label := DomainLabelFromBytes(buf)
		d.Labels = append(d.Labels, *label)
		buf = buf[label.Length + 1:]
	}
	return &d
}

func QuestionFromBytes(buf []byte) *Question {
	q := Question{}
	q.Name = *DomainNameFromBytes(buf)
	buf = buf[len(q.Name.Bytes())+1:]
	q.Type = RecordType(binary.BigEndian.Uint16(buf[:2]))
	q.Class = RecordClass(binary.BigEndian.Uint16(buf[2:4]))
	return &q
}

func ResourceRecordFromBytes(buf []byte) *ResourceRecord {
	r := ResourceRecord{}
	r.Name = *DomainNameFromBytes(buf)
	buf = buf[len(r.Name.Bytes())+1:]
	r.Type = RecordType(binary.BigEndian.Uint16(buf[:2]))
	r.Class = RecordClass(binary.BigEndian.Uint16(buf[2:4]))
	r.TTL = binary.BigEndian.Uint32(buf[4:8])
	r.RDLength = binary.BigEndian.Uint16(buf[8:10])
	r.RData = buf[10:int(r.RDLength)]
	return &r
}

func MessageFromBytes(buf []byte) *Message {
	m := Message{}
	m.Header = *HeaderFromBytes(buf)
	buf = buf[12:]
	m.Questions = make([]Question, m.Header.QDCount)
	for i := 0; i < int(m.Header.QDCount); i++ {
		q := QuestionFromBytes(buf)
		m.Questions = append(m.Questions, *q)
		buf = buf[len(q.Bytes()):]
	}
	m.Answers = make([]ResourceRecord, m.Header.ANCount)
	for i := 0; i < int(m.Header.ANCount); i++ {
		r := ResourceRecordFromBytes(buf)
		m.Answers = append(m.Answers, *r)
		buf = buf[len(r.Bytes()):]
	}
	return &m
}
