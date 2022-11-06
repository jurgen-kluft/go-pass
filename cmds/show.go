package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jurgen-kluft/go-pass/closestmatch"
	"github.com/jurgen-kluft/go-pass/qrterminal"
	"github.com/jurgen-kluft/go-pass/repo"
)

// - show [ --line[=line-number or field-name], -l[line-number or field-name] ] [--qrcode[=line-number or field-name], -q[line-number or field-name] ] pass-name

type ShowCmd struct {
	LineNumbers bool    `optional:"" short:"l" help:"Show the content with line numbers"`
	Line        *int    `optional:"" short:"i" help:"Show the content of a specific line"`
	Field       *string `optional:"" short:"f" help:"Show the content of a specific field"`
	QrCode      *string `optional:"" short:"q" help:"Show the content of a specific field as a QR code (e.g. -q \"site\")"`
	PassName    string  `arg:"" help:"Pass name."`
}

func (a *ShowCmd) Run(globals *Globals) error {
	r := &repo.Repo{}

	r.Root = "$HOME/Documents/Vault"
	r.Root = os.ExpandEnv(r.Root)

	r.Scan()

	fmt.Println(r.Root)

	qrcode := ""
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

					if a.QrCode != nil {
						if strings.HasPrefix(string(line), *a.QrCode) {
							qrcode = string(line)
						}
					}

					if a.LineNumbers {
						fmt.Printf("%d: %s\n", lineNumber, string(line))
					} else if a.Line != nil && *a.Line == lineNumber {
						fmt.Printf("%s\n", string(line))
					} else {
						fmt.Printf("%s\n", string(line))
					}
				})
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}
			}
		}
	}

	if a.QrCode != nil && qrcode != "" {
		// Generate a 'dense' qrcode with the 'Low' level error correction and write it to Stdout
		qrcode = qrcode[len(*a.QrCode)+1:]
		qrcode = strings.TrimSpace(qrcode)

		fmt.Printf("QR code: %s\n", qrcode)
		qrterminal.Generate(qrcode, qrterminal.L, os.Stdout)
	}

	return nil
}
