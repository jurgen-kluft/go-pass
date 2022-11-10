package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jurgen-kluft/go-pass/closestmatch"
	"github.com/jurgen-kluft/go-pass/repo"
)

/*
List names of passwords inside the tree that match pass-names by using the tree(1) program.

This command is alternatively named search.

e.g. gpw search -c email sophia
*/
type SearchCmd struct {
	Fuzzy   int    `short:"f" long:"fuzzy" default:"10" help:"Number of closest matches to show."`
	Search  string `arg:"" help:"Pass name to search."`
	Content string `arg:"" optional:"" help:"Content to search."`
}

func (a *SearchCmd) Run(globals *Globals) error {
	r := &repo.Repo{Root: globals.Root}
	r.Root = os.ExpandEnv(r.Root)
	if err := r.Scan(); err != nil {
		return err
	}

	if a.Search != "" {
		wordsToTest := []string{}
		for _, gr := range r.Files {
			for _, f := range gr {
				wordsToTest = append(wordsToTest, f)
			}
		}
		bagSizes := []int{2}                          // Choose a set of bag sizes, more is more accurate but slower
		cm := closestmatch.New(wordsToTest, bagSizes) // Create a closestmatch object

		// List all the pass-names of the subfolder
		closest := cm.ClosestN(a.Search, a.Fuzzy)
		for _, f := range closest {
			if len(a.Content) > 0 {

				filename := filepath.Join(r.Root, f)

				matches, err := repo.SearchInFile(filename, func(line []byte) error {
					if strings.Contains(string(line), a.Content) {
						return repo.MatchError
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

			} else {
				println(f)
			}
		}
	}
	return nil
}
