package main

import (
	"fmt"
	"os"

	kong "github.com/alecthomas/kong"
)

/*
pass --list

   list all the groups and their items

commands:

- init [ --path=sub-folder, -p sub-folder ]
- set-env [ pass-name ]
- insert FILENAME pass-name
- show [ --clip[=line-number or field-name], -c[line-number or field-name] ] [--qrcode[=line-number or field-name], -q[line-number or field-name] ] pass-name
- search pass-names
- grep [GREPOPTIONS] search-string
- remove [ --recursive, -r ] [ --force, -f ] pass-name
- rename [ --force, -f ] old-path new-path
- copy [ --force, -f ] old-path new-path
- qr pass-name [field-name (e.g. 'site', 'email')]
*/

type InitCmd struct {
	SubFolder string `arg:"" short:"p" help:"Sub folder."`
}

type SetEnvCmd struct {
	PassName string `arg:"" short:"p" help:"Pass name filter."`
}

type InsertCmd struct {
	Filename *os.File `short:"f" help:"File to load content from to post as new entry."`
	PassName string   `short:"p" help:"Pass name (if not provided, name of file is used)."`
}

type Globals struct {
	Config  string      `help:"Location of client config files" default:"~/.docker" type:"path"`
	Version VersionFlag `name:"version" help:"Print version information and quit"`
}

type CLI struct {
	Globals

	List   ListCmd   `cmd:"" default:"withargs" help:"List all pass entries."`
	Init   InitCmd   `cmd:"" help:"Initialize a new pass repository."`
	SetEnv SetEnvCmd `cmd:"" help:"For any pass entry that has 'env: NAME' actually export an environment variable 'NAME = {field}."`
	Insert InsertCmd `cmd:"" help:""`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("pass"),
		kong.Description("managing passwords"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
