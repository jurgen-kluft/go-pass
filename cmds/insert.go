package cmds

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// - insert [ FILE ] pass-name
// e.g. gpw insert sites/www.news.com

type InsertCmd struct {
	File     *os.File `optional:"" help:"File to load content from to post as new entry."`
	PassName string   `arg:"" help:"Pass name of entry to create."`
}

func (c *InsertCmd) Run(globals *Globals) error {
	if c.File != nil {
		// create a new file using c.PassName
		// and write the content of c.File to it
		newfilename := filepath.Join(globals.Root, c.PassName)
		newfile, err := os.Create(newfilename)
		if err != nil {
			return err
		}
		_, err = io.Copy(newfile, c.File)
		c.File.Close()
		newfile.Close()

	} else {
		var site string
		fmt.Print("site: ")
		fmt.Scanln(&site)
	}

	return nil
}
