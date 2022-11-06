package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jurgen-kluft/go-pass/closestmatch"
	"github.com/jurgen-kluft/go-pass/repo"
)

// - show [ --line[=line-number or field-name], -l[line-number or field-name] ] [--qrcode[=line-number or field-name], -q[line-number or field-name] ] pass-name

type ShowCmd struct {
	PassName string `arg:"" help:"Pass name."`
}

func (a *ShowCmd) Run(globals *Globals) error {
	r := &repo.Repo{}

	r.Root = "$HOME/Documents/Vault"
	r.Root = os.ExpandEnv(r.Root)

	r.Scan()

	fmt.Println(r.Root)

	if a.PassName != "" {
		wordsToTest := []string{}
		for _, gr := range r.Files {
			for _, f := range gr {
				wordsToTest = append(wordsToTest, f)
			}
		}
		bagSizes := []int{2}                          // Choose a set of bag sizes, more is more accurate but slower
		cm := closestmatch.New(wordsToTest, bagSizes) // Create a closestmatch object

		// List the content of a pass-name
		filename := cm.Closest(a.PassName)
		if len(filename) > 0 {
			fullFilename := filepath.Join(r.Root, filename)
			if _, err := os.Stat(fullFilename); os.IsNotExist(err) {
				return errors.New(fmt.Sprintf("File '%s' does not exist", filename))
			} else {
				// print name of file
				fmt.Printf("Pass '%s'\n", filename)
				fmt.Println()

				err := repo.GetFileContentByLine(fullFilename, func(lineNumber int, line []byte) {
					fmt.Printf("%d: %s\n", lineNumber, string(line))
				})
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}
			}
		}
	}
	return nil
}
