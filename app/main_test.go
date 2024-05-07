package main

import (
	"bytes"
	"os"
	"path"
	"testing"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

func TestHandleRequest(t *testing.T) {
	var res dns.Message
	tcs := []struct {
		n string
		data []byte
		expected  []byte
	}{
		{
			n: "test handleRequest",
			data: []byte{142,195,1,0,0,1,0,0,0,0,0,0,12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0,0,1,0,1},
			expected: res.Bytes(),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			actual := handleRequest(tc.data, "8.8.8.8:53")
			if !bytes.Equal(actual, tc.expected){
				t.Errorf("Expected %v, got %v", string(tc.expected), string(actual))
			}
		})
	}	
}

func cdProjectRoot(t *testing.T) {
	t.Helper()
	d, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current directory: %s", err)
	}
	findProjectRoot(t, d)
}

func findProjectRoot(t *testing.T, dir string) {
	t.Logf("Looking for .git in %s", dir)
	fp := path.Join(dir, ".git")

	if _, err := os.Stat(fp); err == nil {
		t.Logf("Found .git in %s", dir)
		os.Chdir(dir)
		return
	}
	findProjectRoot(t, dir + "/..")
}