package cmds

import (
	"github.com/jurgen-kluft/go-pass/repo"
)

type ListCmd struct {
	Group    string `short:"g" help:"Group filter."`
	PassName string `short:"p" help:"Pass filter."`
}

func (a *ListCmd) Run(globals *Globals) error {
	r := &repo.Repo{}
	r.Root = "/Users/obnosis5/Documents/Vault"
	r.Scan()
	return nil
}
