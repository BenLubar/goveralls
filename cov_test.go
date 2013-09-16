package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseCov(t *testing.T) {
	cov, err := ioutil.ReadFile("goveralls-test/coverage.json")
	if err != nil {
		t.Fatal(err)
		return
	}

	files := ParseCov(cov, "/home/ben/go/src/github.com/BenLubar/goveralls")
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
