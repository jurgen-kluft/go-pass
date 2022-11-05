package cmds

// - search pass-names

type SearchCmd struct {
	PassName string `short:"p" help:"Pass name (if not provided, name of file is used)."`
}
