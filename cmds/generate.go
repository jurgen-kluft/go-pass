package cmds

import "fmt"

type GenerateCmd struct {
	NoSymbols bool   `optional:"" short:"n" name:"no-symbols" help:"Password should not contain symbols"`
	Clip      bool   `optional:"" short:"c" help:"Copy generated password to clipboard"`
	InPlace   bool   `optional:"" short:"i" name:"in-place" help:"Set generated password as new password for 'pass-name'"`
	Form      string `optional:"" short:"f" help:"The form to use, e.g. 'xxxx-xxxx-xxxxxxxx-xxxx-xxxx'" default:"xxxx-xxxx-xxxxxxxx-xxxx-xxxx"`
	Length    int    `optional:"" short:"l" help:"The length of the password to generate"`
	PassName  string `arg:"" help:"If provided, set the generated password as the new password for 'pass-name'"`
}

func (a *GenerateCmd) Run(globals *Globals) error {
	fmt.Printf("No Symbols: %v\n", a.NoSymbols)
	fmt.Printf("Copy to clipboard: %v\n", a.Clip)
	fmt.Printf("In place: %v\n", a.InPlace)
	fmt.Printf("Form: %s\n", a.Form)
	fmt.Printf("Length: %d\n", a.Length)
	fmt.Printf("Pass name: %s\n", a.PassName)
	return nil
}
