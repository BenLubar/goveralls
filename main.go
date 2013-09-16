// goveralls is a Go client for coveralls.io.
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
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

	cmd := exec.Command("git", "log", "--format=%H", "-n", "1", "HEAD")
	cmd.Stderr = os.Stderr
	revision, err := cmd.Output()
	if err != nil {
		log.Fatalf("Git error: %v", err)
	}
	job.ID = string(revision)

	cmd = exec.Command("gocov", "test", "./...")
	cmd.Stderr = os.Stderr
	cov, err := cmd.Output()
	if err != nil {
		log.Fatalf("gocov error: %v", err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	job.Files = ParseCov(cov, wd)
	Submit(&job)
}
