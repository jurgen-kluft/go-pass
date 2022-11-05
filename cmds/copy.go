package cmds

import "fmt"

// - copy [ --force, -f ] old-path new-path

type CopyCmd struct {
	Force   bool   `short:"f" help:"Force copy even if 'new-path' exists"`
	OldPath string `arg:"" help:"Old pass name to use as source."`
	NewPath string `arg:"" help:"New pass name used as destination."`
}

func (c *CopyCmd) Run(globals *Globals) error {
	fmt.Printf("Force: %v\n", c.Force)
	fmt.Printf("Old path: %s\n", c.OldPath)
	fmt.Printf("New path: %s\n", c.NewPath)
	return nil
}
