package cmds

import (
	"fmt"
	"os"
)

// - insert FILE pass-name

type InsertCmd struct {
	File     *os.File `arg:"" help:"File to load content from to post as new entry."`
	PassName string   `arg:"" help:"Pass name (if not provided, name of file is used)."`
}

func (c *InsertCmd) Run(globals *Globals) error {
	fmt.Printf("File: %v\n", c.File)
	fmt.Printf("Pass name: %s\n", c.PassName)
	return nil
}
