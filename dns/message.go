package dns

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

type Authority struct {
}

type Additional struct {
}

type Message struct {
    Header      Header           `json:"header"`
    Questions   []Question       `json:"questions"`
    Answers     []ResourceRecord `json:"answers"`
    Authorities []Authority      `json:"authorities"`
    Additionals []Additional     `json:"additionals"`
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

func MessageFromBytes(buf []byte) *Message {
	m := Message{}
	m.Header = HeaderFromBytes(buf)
	qbuf := buf[12:]
	
	for i := 0; i < int(m.Header.QDCount); i++ {
		q := QuestionFromBytes(qbuf, buf)
		m.Questions = append(m.Questions, q)
		o := len(q.Bytes())
		if o > len(qbuf) {
			qbuf = qbuf[5:]
		} else {
			qbuf = qbuf[o:]
		}
		
	}
	
	for i := 0; i < int(m.Header.ANCount); i++ {
		r := ResourceRecordFromBytes(buf, buf)
		m.Answers = append(m.Answers, r)
		buf = buf[len(r.Bytes()):]
	}
	return &m
}
