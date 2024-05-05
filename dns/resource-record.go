package dns

import (
	"encoding/binary"
	"fmt"
)

type ResourceRecord struct {
    Name     DomainName  `json:"name"`
    Type     RecordType  `json:"type"`
    Class    RecordClass `json:"class"`
    TTL      uint32      `json:"ttl"`
    RDLength uint16      `json:"rdlength"`
    RData    []byte      `json:"rdata"`
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

func ResourceRecordFromBytes(buf []byte) ResourceRecord {
	r := ResourceRecord{}
	r.Name = DomainNameFromBytes(buf)
	buf = buf[len(r.Name.Bytes()):]
	r.Type = RecordType(binary.BigEndian.Uint16(buf[:2]))
	r.Class = RecordClass(binary.BigEndian.Uint16(buf[2:4]))
	r.TTL = binary.BigEndian.Uint32(buf[4:8])
	r.RDLength = binary.BigEndian.Uint16(buf[8:10])
	r.RData = buf[10:int(r.RDLength)]
	return r
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