package cmds

// - set-env [ pass-name ]

type SetEnvCmd struct {
	PassName string `arg:"" short:"p" help:"Pass name filter."`
}
