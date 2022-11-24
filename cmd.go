package pager

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

/*
func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
}
*/

var Cmd = &Z.Cmd{
	Name: `page`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Description: `
		The {{aka}} command detects if the current runtime environment
		has a pager program of any kind (like {{cmd "less"}} or
		{{cmd "more"}}) and simply executes that instead, but if it cannot
		find a pager it provides its own rudimentary one that attempts to
		read the dimensions of the terminal and properly page content.

		{{cmd .Name}} depends on the {{pkg "term"}} package for paging.

		Avoid temptation to rename this to {{exe "pager"}} since that is
		a reserved executable name on all UNIX/Linux systems.

	`,
}
