package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

var fwdAddr = flag.String("resolver", "", "Forward DNS resolver address")

func main() {

	flag.Parse()
	
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
		
	
		_, err = udpConn.WriteToUDP(handleRequest(buf[:size], *fwdAddr), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}

func handleRequest(data []byte, fwdAddr string) []byte {
	fmt.Println(data)
	fmt.Println(fwdAddr)

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

	if fwdAddr != "" {
		// for each question in the request, send a query to the forward resolver
		// and append the response to the response message
		for _, q := range req.Questions {
			msg := dns.Message{}
			msg.Header = dns.Header{
				ID: req.Header.ID,
				QR: req.Header.QR,
				AA: req.Header.AA,
				TC: req.Header.TC,
				RD: req.Header.RD,
				RA: req.Header.AA,
				Z: req.Header.Z,
				Opcode: req.Header.Opcode,
				RCode: req.Header.RCode,
				QDCount: 1,
				ANCount: 0,
				NSCount: 0,
				ARCount: 0,
			}
			msg.Questions = append(msg.Questions, q)
			msgBuf := msg.Bytes()
			resBuf := handleFwdRequest(msgBuf, fwdAddr)
			resMsg := dns.MessageFromBytes(resBuf)
			res.Answers = append(res.Answers, resMsg.Answers...)

		}
	} else {

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
		
		
	}
	res.Header.ANCount = uint16(len(res.Answers))
	return res.Bytes()
}

func handleFwdRequest(data []byte, fwdAddr string) []byte {
	conn, err := net.Dial("udp", fwdAddr)
	if err != nil {
		fmt.Println("Failed to connect to forward resolver:", err)
		return nil
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Failed to send data to forward resolver:", err)
		return nil
	}
	buf := make([]byte, 1024)
	size, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read data from forward resolver:", err)
		return nil
	}
	return buf[:size]

}
