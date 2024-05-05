package dns

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"
)

func loadJson(t *testing.T, fn string, sn string) any {
	t.Helper()
	cdProjectRoot(t)
	fp := path.Join("dns", "testdata", fmt.Sprintf("%s.json", fn))
	fmt.Println("Loading file", fp)
	data, err := os.ReadFile(fp)
	if err != nil {
		t.Fatalf("Could not load file %s", fn)
	}

	switch sn {
	case "header":
		var obj Header
		json.Unmarshal(data, &obj)
		return obj
	case "resource-record":
		var obj ResourceRecord
		json.Unmarshal(data, &obj)
		return obj
	case "message":
		var obj Message
		json.Unmarshal(data, &obj)
		return obj
	case "question":
		var obj Question
		json.Unmarshal(data, &obj)
		return obj
	default:
		t.Fatalf("Unknown file %s", fn)
		os.Exit(1)
	}
	return nil
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