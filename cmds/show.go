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
	LineNumbers bool   `optional:"" short:"l" help:"Show the content with line numbers"`
	Line        int    `optional:"" short:"i" help:"Show the content of a specific line"`
	Field       string `optional:"" short:"f" help:"Show the content of a specific field"`
	QrCode      bool   `optional:"" short:"q" help:"Show the content as a QR code"`
	PassName    string `arg:"" help:"Pass name."`
}

func (a *ShowCmd) Run(globals *Globals) error {
	r := &repo.Repo{}

	r.Root = "$HOME/Documents/Vault"
	r.Root = os.ExpandEnv(r.Root)

	r.Scan()

	fmt.Println(r.Root)

	qrcode := ""
	fullQrCode := ""
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

				err := repo.GetFileContentByLine(fullFilename, func(lineNumber int, line []byte) error {
					if a.QrCode {
						if lineNumber == a.Line {
							qrcode = string(line)
						}
						if len(a.Field) > 0 {
							if strings.HasPrefix(string(line), a.Field) {
								qrcode = string(line)
							}
						}
						fullQrCode = fullQrCode + " / " + string(line)
					}

					if a.LineNumbers {
						fmt.Printf("%d: %s\n", lineNumber, string(line))
					} else if a.Line >= 0 {
						if a.Line == lineNumber {
							fmt.Printf("%s\n", string(line))
						}
					} else {
						fmt.Printf("%s\n", string(line))
					}

					return nil
				})
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}
			}
		}
	}

	if a.QrCode {
		if len(qrcode) > 0 {
			// Generate a 'dense' qrcode with the 'Low' level error correction and write it to Stdout
			qrcode = qrcode[len(*&a.Field):]
			qrcode = strings.TrimSpace(qrcode)

			// Check if QR code strings starts with the name of a field followed by a ':', e.g. 'name:'
			// If so remove that (prefix) field from the QR code string
			if strings.Contains(qrcode, ":") {
				field := qrcode[:strings.Index(qrcode, ":")+1]
				if strings.HasPrefix(qrcode, field) {
					qrcode = qrcode[len(field):]
				}
			}
			qrterminal.Generate(qrcode, qrterminal.L, os.Stdout)
		} else if len(fullQrCode) > 0 {
			qrterminal.Generate(fullQrCode, qrterminal.L, os.Stdout)
		}
	}

	return nil
}
