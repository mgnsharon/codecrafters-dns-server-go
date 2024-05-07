package dns

import "encoding/binary"

type DomainLabel struct {
    Length  uint8
    Content string
}

type DomainName struct {
    Labels []DomainLabel
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

func DomainLabelFromBytes(buf []byte) DomainLabel {
	d := DomainLabel{}
	d.Length = buf[0]
	d.Content = string(buf[1:d.Length + 1])
	return d
}

func parseLabels(buf []byte, msgBuf []byte, labels []DomainLabel) ([]byte, []DomainLabel) {
	if buf[0] == 0 {
		return buf, labels
	}
	if buf[0] & 0xC0 == 0xC0{
		offset := binary.BigEndian.Uint16([]byte{buf[0] & 0x3F, buf[1]})
		return parseLabels(msgBuf[offset:], msgBuf, labels)
	}
	d := DomainLabel{}
	d.Length = buf[0]
	d.Content = string(buf[1:d.Length + 1])
	labels = append(labels, d)
	return parseLabels(buf[d.Length + 1:], msgBuf, labels)
}

func DomainNameFromBytes(buf []byte, msgBuf []byte) DomainName {
	d := DomainName{}
	//d.Labels = make([]DomainLabel, 0)
	dnBuf := buf[:]
	for {
		if dnBuf[0] == 0 {
			break
		}
		b, labels := parseLabels(buf, msgBuf, []DomainLabel{})
		d.Labels = append(d.Labels, labels...)
		dnBuf = b[:]
		
	}
	return d
}