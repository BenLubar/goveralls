package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Job struct {
	ID      string  `json:"service_job_id"`
	Service string  `json:"service_name"`
	Files   []*File `json:"source_files"`
}

type File struct {
	Name     string   `json:"name"`
	Source   string   `json:"source"`
	Coverage []*int64 `json:"coverage"`

	offsets []int `json:"-"`
}

func Submit(job *Job) {
	b, err := json.Marshal(job)
	if err != nil {
		panic(err)
	}
	resp, err := http.PostForm("https://coveralls.io/api/v1/jobs", url.Values{
		"repo_token": {*repo_token},
		"json_file":  {string(b)},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)
		return
	}
	io.Copy(os.Stderr, resp.Body)
	os.Exit(1)
}
