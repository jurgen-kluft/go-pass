package repo

import (
	"bufio"
	"bytes"
	"errors"
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

	// Start with a new and empty map
	r.Files = make(map[string][]string)

	// check if the directory in r.Root exists
	if _, err := os.Stat(r.Root); os.IsNotExist(err) {
		return errors.New("Root directory does not exist")
	}

	n := len(r.Root) + 1

	filepath.Walk(r.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if info.IsDir() {
			// hidden directories to be ignored
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
		} else {
			fp := path[n:]
			dir := filepath.Dir(fp)

			// files to ignore
			if info.Name() == ".DS_Store" {
				return nil
			}

			// directories to include (rest is ignored)
			if dir == "people" || dir == "email" || dir == "phone" || dir == "sites" {
				if _, ok := r.Files[dir]; !ok {
					r.Files[dir] = []string{}
				}
				r.Files[dir] = append(r.Files[dir], fp)
			}
		}
		return nil
	})
	return nil
}

// Search the repo using 'glob' patterns
func (r *Repo) GlobSearch(query string, seperators ...rune) ([]FileMatch, error) {
	g, err := glob.Compile(query, seperators...)
	if err != nil {
		return nil, err
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
	return matches, nil
}

func (r *Repo) DirectSearch(query string) ([]FileMatch, error) {
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
	return matches, nil
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
func (r *Repo) GrepSearch(query string) ([]FileMatch, error) {
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
		return matches, nil
	}
	return nil, err
}

func GetFileContentByLine(filename string, handleLine func(lineNumber int, line []byte)) error {

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		text := scanner.Bytes()
		if line == 1 {
			if bytes.IndexByte(text, 0) != -1 {
				return errors.New("cannot show binary file")
			}
		}
		handleLine(line, text)
		line++
	}
	return nil
}

func SearchInFile(filename string, compare func([]byte) bool) ([]Match, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	linematches := []Match{}
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
