package dns

import "encoding/binary"

// Header represents the DNS message header.
type Header struct {
    ID      uint16 `json:"id"` // ID is the identification number of the DNS message.
    QR      uint8  `json:"qr"` // QR is the query/response flag.
    Opcode  uint8  `json:"opcode"` // Opcode is the operation code.
    AA      uint8  `json:"aa"` // AA is the authoritative answer flag.
    TC      uint8  `json:"tc"` // TC is the truncation flag.
    RD      uint8  `json:"rd"` // RD is the recursion desired flag.
    RA      uint8  `json:"ra"` // RA is the recursion available flag.
    Z       uint8  `json:"z"` // Z is reserved for future use.
    RCode   uint8  `json:"rcode"` // RCode is the response code.
    QDCount uint16 `json:"qdcount"` // QDCount is the number of questions in the question section.
    ANCount uint16 `json:"ancount"` // ANCount is the number of resource records in the answer section.
    NSCount uint16 `json:"nscount"` // NSCount is the number of name server resource records in the authority section.
    ARCount uint16 `json:"arcount"` // ARCount is the number of resource records in the additional section.
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

func HeaderFromBytes(buf []byte) Header {
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
	return h
}