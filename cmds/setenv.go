package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jurgen-kluft/go-pass/repo"
)

// - set-env [ pass-name ]

type SetEnvCmd struct {
	PassName string `arg:"" optional:"" help:"Pass name."`
}

func (a *SetEnvCmd) Run(globals *Globals) error {
	r := &repo.Repo{Root: globals.Root}
	r.Root = os.ExpandEnv(r.Root)
	if err := r.Scan(); err != nil {
		return err
	}

	for _, gr := range r.Files {
		for _, f := range gr {
			if f == a.PassName || len(a.PassName) == 0 {

				pw := ""
				err := repo.GetFileContentByLine(filepath.Join(r.Root, f), func(lineNumber int, line []byte) error {
					text := string(line)
					if lineNumber == 1 {
						pw = text
					} else {
						if strings.HasPrefix(text, "env:") {
							key := strings.TrimPrefix(text, "env:")
							key = strings.TrimSpace(key)
							key = strings.ToUpper(key)

							fmt.Printf("%s=%s\n", key, pw)
							os.Setenv(key, pw)
						}
					}
					return nil
				})

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
