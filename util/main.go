package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

func main() {
    var testQuestion dns.Question = dns.Question{
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
	}

    data, err := json.Marshal(testQuestion)
    if err != nil {
        log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
    }

    err = os.WriteFile("./dns/testdata/question.json", data, 0644)
    if err != nil {
        log.Fatalf("Error occurred during writing file. Error: %s", err.Error())
    }
}