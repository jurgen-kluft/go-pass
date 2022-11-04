package main

import "fmt"

type ListCmd struct {
	Group    string `short:"g" help:"Group filter."`
	PassName string `short:"p" help:"Pass filter."`
}

func (a *ListCmd) Run(globals *Globals) error {
	fmt.Printf("List: %s\n", "")
	return nil
}
