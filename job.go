package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Job struct {
	ID      string  `json:"service_job_id"`
	Service string  `json:"service_name"`
	Token   string  `json:"repo_token"`
	Files   []*File `json:"source_files"`

	Git struct {
		Head struct {
			ID             string `json:"id"`
			AuthorName     string `json:"author_name"`
			AuthorEmail    string `json:"author_email"`
			CommitterName  string `json:"committer_name"`
			CommitterEmail string `json:"committer_email"`
			Message        string `json:"message"`
		} `json:"head"`

		Branch string `json:"branch"`

		Remotes []*GitRemote `json:"remotes"`
	} `json:"git"`

	RunAt time.Time `json:"run_at"`
}

type File struct {
	Name     string   `json:"name"`
	Source   string   `json:"source"`
	Coverage []*int64 `json:"coverage"`

	offsets []int `json:"-"`
}

type GitRemote struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func Submit(job *Job) {
	b, err := json.Marshal(job)
	if err != nil {
		panic(err)
	}
	resp, err := http.PostForm("https://coveralls.io/api/v1/jobs", url.Values{
		"json": {string(b)},
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
