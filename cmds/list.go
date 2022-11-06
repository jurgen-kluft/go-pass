package cmds

import (
	"github.com/jurgen-kluft/go-pass/closestmatch"
	"github.com/jurgen-kluft/go-pass/repo"
)

/*
list subfolder

List names of passwords inside the tree at subfolder by using the tree(1) program.
This command is alternatively named list.

*/

type ListCmd struct {
	SubFolder string `arg:"" optional:"" help:"Subfolder to list the entries of."`
}

func (a *ListCmd) Run(globals *Globals) error {
	r := &repo.Repo{}

	r.Root = "/Users/obnosis5/Documents/Vault"
	r.Scan()

	if a.SubFolder == "" {
		for _, gr := range r.Files {
			for _, f := range gr {
				println(f)
			}
		}
	} else {
		wordsToTest := []string{}
		for gn := range r.Files {
			wordsToTest = append(wordsToTest, gn)
		}
		bagSizes := []int{2}                          // Choose a set of bag sizes, more is more accurate but slower
		cm := closestmatch.New(wordsToTest, bagSizes) // Create a closestmatch object

		// List all the pass-names of the subfolder
		closest := cm.Closest(a.SubFolder)
		for _, f := range r.Files[closest] {
			println(f)
		}
	}
	return nil
}
