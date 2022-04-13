package templates

import "embed"

//go:embed account
//go:embed discuss
//go:embed document.html
//go:embed index.html
//go:embed succeed.html
//go:embed failed.html
//go:embed ban.html
var FS embed.FS
