package cmds

// - qr pass-name [field-name (e.g. 'site', 'email')]

type QrCodeCmd struct {
	PassName string `short:"p" help:"Pass name (if not provided, name of file is used)."`
}
