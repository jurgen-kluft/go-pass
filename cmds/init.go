package cmds

import "fmt"

// - init sub-folder
/*
init sub-folder

Initialize new password storage

*/
type InitCmd struct {
	SubFolder string `arg:"" short:"p" help:"Sub folder."`
}

func (c *InitCmd) Run(globals *Globals) error {
	fmt.Printf("Folder: %s\n", c.SubFolder)
	return nil
}
