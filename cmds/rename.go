package cmds

import "fmt"

/*
mv [ --force, -f ] old-path new-path

Renames the password or directory named old-path to new-path.

This command is alternatively named rename. If --force is specified, silently
overwrite new-path if it exists. If new-path ends in a trailing /, it is always
treated as a directory.

*/
type RenameCmd struct {
	Force   bool   `short:"f" help:"Silently overwrite new-path if it exists"`
	OldPath string `arg:"" help:"The old-path to be renamed."`
	NewPath string `arg:"" help:"The new-path to give to old-path."`
}

func (c *RenameCmd) Run(globals *Globals) error {
	fmt.Printf("Force: %v\n", c.Force)
	fmt.Printf("Old path: %s\n", c.OldPath)
	fmt.Printf("New path: %s\n", c.NewPath)
	return nil
}
