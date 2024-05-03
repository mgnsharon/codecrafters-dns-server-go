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
	
		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)
	
		// Create an empty response
		response := dns.Message{
			Header: dns.Header{
				ID: 1234,
				QR: 1,
				Opcode: 0,
				AA: 0,
				TC: 0,
				RD: 0,
				RA: 0,
				Z: 0,
				RCode: 0,
				QDCount: 1,
				ANCount: 1,
				NSCount: 0,
				ARCount: 0,
			},
			Questions: []dns.Question{
				{
					Name: dns.DomainName{
						Labels: []dns.DomainLabel{
							{
								Length: 12,
								Content: "codecrafters",
							},
							{
								Length: 2,
								Content: "io",
							},
						},
					},
					Type: dns.A,
					Class: dns.IN,
				},
								
			},
			Answers: []dns.ResourceRecord{
				{
					Name: dns.DomainName{
						Labels: []dns.DomainLabel{
							{
								Length: 12,
								Content: "codecrafters",
							},
							{
								Length: 2,
								Content: "io",
							},
						},
					},
					Type: dns.A,
					Class: dns.IN,
					TTL: 60,
					RData: (*dns.IPv4Address)(&dns.IPv4Address{
							Octets: [4]uint8{127, 0, 0, 1},
						}).Bytes(),
					RDLength: 4,
				},
			},
		}
			
	
		_, err = udpConn.WriteToUDP(response.Bytes(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
