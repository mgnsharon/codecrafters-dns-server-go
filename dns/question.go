package dns

import "encoding/binary"

type Question struct {
    Name  DomainName  `json:"name"`
    Type  RecordType  `json:"type"`
    Class RecordClass `json:"class"`
}

func (q *Question) Bytes() []byte {
	buf := make([]byte, 0)
	buf = append(buf, q.Name.Bytes()...)
	buf = binary.BigEndian.AppendUint16(buf, uint16(q.Type))
	buf = binary.BigEndian.AppendUint16(buf, uint16(q.Class))
	return buf
}

func QuestionFromBytes(buf []byte, msgBuf []byte) Question {
	q := Question{}
	q.Name = DomainNameFromBytes(buf, msgBuf)
	o := len(q.Name.Bytes())
	if o > len(buf) {
		buf = buf[5:]
	} else {
		buf = buf[o:]
	}
	q.Type = RecordType(binary.BigEndian.Uint16(buf[:2]))
	q.Class = RecordClass(binary.BigEndian.Uint16(buf[2:4]))
	return q
}