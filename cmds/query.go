package cmds

// Show the age of all password-store entries, by looking at the git log.

/*
gpw query [query]

Show the age of all email password-store entries older than 1 year, by looking at the git log.
gpw query "name = email AND age >= 1y"

Possible variables:

- Name
- Content, Line contains a matching string
- Age


	https://github.com/maja42/goval
*/

type QueryCmd struct {
	Query string `arg:"" optional:"" help:"The expression to evaluate to list all matching pass entries"`
}

func (a *QueryCmd) Run(globals *Globals) error {

	// Get the list of all password-store entries and if necessary match them against the pass-name
	// With the list query Git for the last commit date and print the age of the entry

	return nil
}
