package cmds

// - grep [GREPOPTIONS] search-string

type GrepCmd struct {
	PassName string `short:"p" help:"Pass name (if not provided, name of file is used)."`
}
