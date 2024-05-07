package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

func main() {
	
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()
	
	buf := make([]byte, 512)
	
	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
		response := handleRequest(buf[:size])
			
	
		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}

func handleRequest(data []byte) []byte {
	req := dns.MessageFromBytes(data)

	res := dns.Message{}
	
	res.Header.ID = req.Header.ID
	res.Header.QR = 1
	res.Header.AA = 0
	res.Header.TC = 0
	res.Header.RA = 0
	res.Header.Z = 0
	res.Header.RD = req.Header.RD
	res.Header.Opcode = req.Header.Opcode
	if req.Header.Opcode == 0 {
		res.Header.RCode = 0 
	} else {
		res.Header.RCode = 4
	}
	res.Header.QDCount = req.Header.QDCount
	res.Header.ANCount = req.Header.QDCount
	res.Header.NSCount = 0
	res.Header.ARCount = 0

	//res.Questions = append(res.Questions, req.Questions[0])
	res.Questions = append(res.Questions, req.Questions...)

	for _, q := range req.Questions {
		res.Answers = append(res.Answers, dns.ResourceRecord{
			Name: q.Name,
			Type: dns.A,
			Class: dns.IN,
			TTL: 60,
			RDLength: 4,
			RData: (*dns.IPv4Address)(&dns.IPv4Address{
				Octets: [4]uint8{8, 8, 8, 8},
			}).Bytes(),
		})
	}
	res.Authorities = make([]dns.Authority, 0)
	res.Additionals = make([]dns.Additional, 0)
	
	return res.Bytes()
}
