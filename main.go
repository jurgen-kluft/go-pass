package main

import (
	kong "github.com/alecthomas/kong"
	cmds "github.com/jurgen-kluft/go-pass/cmds"
)

const (
	VERSION = "0.1.1"
)

type CLI struct {
	cmds.Globals

	Init     cmds.InitCmd     `cmd:"" help:"Initialize a new pass repository."`
	List     cmds.ListCmd     `cmd:"" help:"List all pass entries."`
	SetEnv   cmds.SetEnvCmd   `cmd:"" help:"For any pass entry that has 'env: NAME' actually export an environment variable 'NAME = {field}."`
	Insert   cmds.InsertCmd   `cmd:"" help:"Insert/Create a new pass entry"`
	Show     cmds.ShowCmd     `cmd:"" help:"Show details/content of an existing pass entry"`
	Search   cmds.SearchCmd   `cmd:"" help:"Search for one or more pass entries"`
	Grep     cmds.GrepCmd     `cmd:"" help:"Search (using grep) for one or more entries"`
	Generate cmds.GenerateCmd `cmd:"" help:"Generate a new password and optionally update it for an existing pass entry"`
	Remove   cmds.RemoveCmd   `cmd:"" help:"Remove/Delete an existing pass entry"`
	Rename   cmds.RenameCmd   `cmd:"" help:"Rename/Move an existing pass entry"`
	Copy     cmds.CopyCmd     `cmd:"" help:"Copy/Duplicate an existing pass entry"`
}

func main() {
	cli := CLI{
		Globals: cmds.Globals{
			Root:    "$HOME/Documents/Vault",
			Version: cmds.VersionFlag(VERSION),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("gpw"),
		kong.Description("managing passwords"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": VERSION,
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
