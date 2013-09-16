package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestParseCov(t *testing.T) {
	cmd := exec.Command("gocov", "test", "github.com/BenLubar/goveralls/goveralls-test")
	cmd.Stderr = os.Stderr
	cov, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}

	files := ParseCov(cov, wd)
	filesJson, err := json.Marshal(files)
	if err != nil {
		t.Fatal(err)
		return
	}

	expected, err := ioutil.ReadFile("goveralls-test/expected.json")
	if err != nil {
		t.Fatal(err)
		return
	}

	if !bytes.Equal(filesJson, expected) {
		t.Errorf("Expected:\t%q", expected)
		t.Errorf("Actual:\t%q", filesJson)
	}
}
