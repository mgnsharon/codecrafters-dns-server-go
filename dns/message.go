package dns

import (
	"encoding/binary"
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

type Question struct {
}

type Answer struct {
}

type Authority struct {
}

type Additional struct {
}

type Message struct {
	Header     Header
	Questions  []Question
	Answers    []Answer
	Authorities []Authority
	Additionals []Additional
}

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
	return buf
}
