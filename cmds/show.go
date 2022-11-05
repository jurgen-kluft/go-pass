package cmds

// - show [ --clip[=line-number or field-name], -c[line-number or field-name] ] [--qrcode[=line-number or field-name], -q[line-number or field-name] ] pass-name

type ShowCmd struct {
	PassName string `short:"p" help:"Pass name (if not provided, name of file is used)."`
}
