package cmds

import (
	"github.com/jurgen-kluft/go-pass/closestmatch"
	"github.com/jurgen-kluft/go-pass/repo"
)

/*
List names of passwords inside the tree that match pass-names by using the tree(1) program.

This command is alternatively named search.
*/
type SearchCmd struct {
	Closest  int    `short:"c" long:"closest" default:"10" help:"Number of closest matches to show."`
	PassName string `arg:"" help:"Pass name to search (closest matches)."`
}

func (a *SearchCmd) Run(globals *Globals) error {
	r := &repo.Repo{}

	r.Root = "/Users/obnosis5/Documents/Vault"
	r.Scan()

	if a.PassName != "" {
		wordsToTest := []string{}
		for _, gr := range r.Files {
			for _, f := range gr {
				wordsToTest = append(wordsToTest, f)
			}
		}
		bagSizes := []int{2}                          // Choose a set of bag sizes, more is more accurate but slower
		cm := closestmatch.New(wordsToTest, bagSizes) // Create a closestmatch object

		// List all the pass-names of the subfolder
		closest := cm.ClosestN(a.PassName, a.Closest)
		for _, f := range closest {
			println(f)
		}
	}
	return nil
}
