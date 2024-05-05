package dns

type DomainLabel struct {
    Length  uint8  `json:"length"`
    Content string `json:"content"`
}

type DomainName struct {
    Labels []DomainLabel `json:"labels"`
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

func DomainNameFromBytes(buf []byte) DomainName {
	d := DomainName{}
	d.Labels = make([]DomainLabel, 0)
	for {
		if buf[0] == 0 {
			break
		}
		label := DomainLabelFromBytes(buf)
		d.Labels = append(d.Labels, label)
		buf = buf[label.Length + 1:]
	}
	return d
}