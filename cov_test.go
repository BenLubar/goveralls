package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseCov(t *testing.T) {
	if output, err := exec.Command("go", "get", "github.com/axw/gocov/gocov").CombinedOutput(); err != nil {
		t.Log(string(output))
		t.Fatal(err)
		return
	}

	path := filepath.SplitList(os.Getenv("PATH"))
	for _, gopath := range filepath.SplitList(os.Getenv("GOPATH")) {
		path = append(path, filepath.Join(gopath, "bin"))
	}
	os.Setenv("PATH", strings.Join(path, string(filepath.ListSeparator)))

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
	if err != nil {
		t.Fatal(err)
		return
	}

	expectedJson, err := ioutil.ReadFile("goveralls-test/expected.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	var expected []*File
	err = json.Unmarshal(expectedJson, &expected)
	if err != nil {
		t.Fatal(err)
		return
	}

	filesJson, _ := json.Marshal(files)
	expectedJson, _ = json.Marshal(expected)
	if !bytes.Equal(filesJson, expectedJson) {
		t.Errorf("Actual:  \t%q", filesJson)
		t.Errorf("Expected:\t%q", expectedJson)
	}
}
