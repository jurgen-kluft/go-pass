package cmds

// - remove [ --recursive, -r ] [ --force, -f ] pass-name

type RemoveCmd struct {
	PassName string `short:"p" help:"Pass name (if not provided, name of file is used)."`
}
