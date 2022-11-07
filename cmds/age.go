package cmds

// Show the age of all password-store entries, by looking at the git log.

/*
gpw age [pass-name]
*/

type AgeCmd struct {
	PassName string `arg:"" optional:"" help:"If provided, only list those entries that match 'pass-name'"`
}

func (a *AgeCmd) Run(globals *Globals) error {

	// Get the list of all password-store entries and if necessary match them against the pass-name
	// With the list query Git for the last commit date and print the age of the entry

	return nil
}
