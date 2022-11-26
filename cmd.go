package page

import (
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:    `page`,
	Summary: `use system pager or backup`,
	MaxArgs: 1,
	Usage:   ` [FILE]`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Description: `
		The {{aka}} command detects if the current runtime environment
		has a page program of any kind (like {{cmd "less"}} or
		{{cmd "more"}}) and simply executes that instead, but if it cannot
		find a page it provides its own rudimentary one that attempts to
		read the dimensions of the terminal and properly page content.

		{{cmd .Name}} accepts a single optional argument for the file to
		page, but usually content is piped into standard input.

		{{cmd .Name}} depends on the {{pkg "term"}} package for paging.

		Avoid temptation to rename this to {{exe "pager"}} since that is
		a reserved executable name on all UNIX/Linux systems.

	`,

	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return This(os.Stdin)
		}
		return File(args[0])
	},
}
