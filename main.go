// goveralls is a Go client for coveralls.io.
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	repo_token = flag.String("repo_token", "", "found at the bottom of your repository's page on coveralls.io.")
)

func main() {
	flag.Parse()
	if *repo_token == "" {
		flag.Usage()
		os.Exit(2)
	}

	var job Job
	job.Service = "goveralls"
	job.Token = *repo_token

	job.Git.Branch = git("rev-parse", "--abbrev-ref", "HEAD")
	job.ID = git("log", "-1", "--format=%H")
	job.Git.Head.ID = job.ID
	job.Git.Head.AuthorName = git("log", "-1", "--format=%aN")
	job.Git.Head.AuthorEmail = git("log", "-1", "--format=%aE")
	job.Git.Head.CommitterName = git("log", "-1", "--format=%cN")
	job.Git.Head.CommitterEmail = git("log", "-1", "--format=%cE")
	job.Git.Head.Message = git("log", "-1", "--format=%s")
	for _, line := range strings.FieldsFunc(git("remote", "-v"), func(r rune) bool { return r == '\n' }) {
		fields := strings.Fields(line)
		job.Git.Remotes = append(job.Git.Remotes, &GitRemote{
			Name: fields[0],
			URL:  fields[1],
		})
	}

	job.RunAt = time.Now()

	cov, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	job.Files = ParseCov(cov, wd)
	Submit(&job)
}

func git(args ...string) string {
	cmd := exec.Command("git", args...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Git error: %v", err)
	}
	return string(output)
}
