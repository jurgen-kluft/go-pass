package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jurgen-kluft/go-pass/repo"
)

// - grep search-string sub-search-string

type GrepCmd struct {
	Content bool   `short:"c" help:"Search in content."`
	BuiltIn string `short:"b" help:"Search in using built-in grep command (e.g. url, email, phone)."`
	Grep    string `arg:"" help:"Use grep to find a matching pass entry."`
	SubGrep string `arg:"" optional:"" help:"Use sub-grep to find a match in the found text."`
}

// a map for built-in grep functions
var builtInGrep = map[string]string{
	"email": "[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+",
	"phone": "[0-9]{3}-[0-9]{3}-[0-9]{4}",
	"url":   "(http|https)://[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+",
}

func (a *GrepCmd) Run(globals *Globals) error {
	r := &repo.Repo{Root: globals.Root}
	r.Root = os.ExpandEnv(r.Root)
	r.Scan()

	if a.Grep != "" {

		rx, err := regexp.Compile(a.Grep)
		if err != nil {
			return err
		}

		var sx *regexp.Regexp
		if len(a.SubGrep) > 0 {
			sx, err = regexp.Compile(a.SubGrep)
			if err != nil {
				return err
			}
		}

		for _, gr := range r.Files {
			for _, f := range gr {

				filename := filepath.Join(r.Root, f)

				matches, err := repo.SearchInFile(filename, func(text []byte) error {
					if rx.Match(text) {
						if sx == nil {
							return repo.MatchError
						} else {
							found := rx.Find(text)
							if sx.Match(found) {
								return repo.MatchError
							}
						}
					}
					return nil
				})
				if err != nil {
					return err
				}
				if len(matches) > 0 {

					fmt.Printf("Matches found in '%s'\n", f)

					for _, m := range matches {
						fmt.Printf("%d: %s\n", m.Line, m.Text)
					}
				}
			}
		}
	}
	return nil
}
