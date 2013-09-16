package main

import (
	"encoding/json"
	"github.com/axw/gocov"
	"io/ioutil"
	"path/filepath"
)

func ParseCov(cov []byte, wd string) []*File {
	var input struct {
		Packages []*gocov.Package
	}
	err := json.Unmarshal(cov, &input)
	if err != nil {
		panic(err)
	}

	files := make(map[string]*File)
	for _, pkg := range input.Packages {
		for _, fun := range pkg.Functions {
			rel, err := filepath.Rel(wd, fun.File)
			if err != nil {
				panic(err)
			}
			file, ok := files[rel]
			if !ok {
				contents, err := ioutil.ReadFile(rel)
				if err != nil {
					panic(err)
				}

				file = &File{
					Name:   rel,
					Source: string(contents),
				}
				for offset, b := range contents {
					if b == '\n' {
						file.offsets = append(file.offsets, offset)
					}
				}
				file.offsets = append(file.offsets, len(contents))
				file.Coverage = make([]*int64, len(file.offsets))

				files[rel] = file
			}

			for _, stmt := range fun.Statements {
				for line, offset := range file.offsets {
					if stmt.Start < offset {
						if file.Coverage[line] == nil {
							file.Coverage[line] = new(int64)
						}
						*file.Coverage[line] += stmt.Reached
						break
					}
				}
			}
		}
	}
	fileSlice := make([]*File, 0, len(files))
	for _, f := range files {
		fileSlice = append(fileSlice, f)
	}
	return fileSlice
}
