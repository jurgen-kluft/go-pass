package repo

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jurgen-kluft/go-pass/glob"
)

type Repo struct {
	Root string

	// List of all the files found
	Files map[string][]string
}

func (r *Repo) Scan() error {
	fmt.Println(r.Root)

	filepath.Walk(r.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
		} else {
			//fmt.Println(path)
			fmt.Printf("File Name: %s\n", info.Name())
		}
		return nil
	})

	return nil
}

// Search the repo using 'glob' patterns
func (r *Repo) GlobSearch(query string, seperators ...rune) []FileMatch {
	g, err := glob.Compile(query, seperators...)
	if err != nil {
		return nil
	}
	matches := []FileMatch{}
	for _, group := range r.Files {
		for _, path := range group {
			m, err := SearchInFile(path, func(line []byte) bool {
				return g.Match(string(line))
			})
			if err == nil {
				matches = append(matches, FileMatch{path, m})
			}
		}
	}
	return matches
}

func (r *Repo) DirectSearch(query string) []FileMatch {
	matches := []FileMatch{}
	for _, group := range r.Files {
		for _, path := range group {
			m, err := SearchInFile(path, func(line []byte) bool {
				return strings.Compare(string(line), query) == 0
			})
			if err == nil {
				matches = append(matches, FileMatch{path, m})
			}
		}
	}
	return matches
}

type Match struct {
	Line int
	Text string
}

type FileMatch struct {
	Path    string
	Matches []Match
}

// Search the repo using 'grep'
func (r *Repo) GrepSearch(query string) []FileMatch {
	reg, err := regexp.Compile(query)
	if err == nil {
		matches := []FileMatch{}
		for _, group := range r.Files {
			for _, path := range group {
				m, err := SearchInFile(path, func(line []byte) bool {
					return reg.Match(line)
				})
				if err == nil {
					matches = append(matches, FileMatch{path, m})
				}
			}
		}
		return matches
	}
	return nil
}

func SearchInFile(filename string, compare func([]byte) bool) (linematches []Match, err error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		text := scanner.Bytes()
		if line == 1 {
			if bytes.IndexByte(text, 0) != -1 {
				return nil, errors.New("cannot search in binary file")
			}
		}
		if compare(text) {
			linematches = append(linematches, Match{line, string(text)})
		}
		line++
	}
	return linematches, nil
}
