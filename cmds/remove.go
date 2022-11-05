package cmds

import "fmt"

/*
remove [ --recursive, -r ] [ --force, -f ] pass-name

Remove the password named pass-name from the password store.

This command is alternatively named remove or delete. If --recursive or -r is
specified, delete pass-name recursively if it is a directory.
If --force or -f is specified, do not interactively prompt before removal.

*/
type RemoveCmd struct {
	Recursive bool   `short:"r" help:"If pass-name is a directory remove all sub entries"`
	Force     bool   `short:"f" help:"Do not prompt for removal, just remove it silently."`
	PassName  string `arg:"" short:"p" help:"Pass name (if not provided, name of file is used)."`
}

func (a *RemoveCmd) Run(globals *Globals) error {
	fmt.Printf("Recursive: %v\n", a.Recursive)
	fmt.Printf("Force: %v\n", a.Force)
	fmt.Printf("Pass name: %v\n", a.PassName)

	return nil
}
