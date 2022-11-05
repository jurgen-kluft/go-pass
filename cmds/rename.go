package cmds

// - rename [ --force, -f ] old-path new-path

type RenameCmd struct {
	PassName string `short:"p" help:"Pass name (if not provided, name of file is used)."`
}
