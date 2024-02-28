package velde

import (
	"context"
	"fmt"
	"io"

	"github.com/youta-t/flarc"
	"github.com/youta-t/velde/velde/subcommands/list"
	"github.com/youta-t/velde/velde/subcommands/uninstall"
)

func Main(
	ctx context.Context,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
	args []string,
) int {

	cmd_list, err := list.Command()
	if err != nil {
		fmt.Fprintf(stderr, "[velde] bug. please report it: %s", err)
		return 1
	}
	cmd_uninstall, err := uninstall.Command()
	if err != nil {
		fmt.Fprintf(stderr, "[velde] bug. please report it: %s", err)
		return 1
	}

	main, err := flarc.NewCommandGroup[struct{}](
		`delve version switcher`,
		struct{}{},
		flarc.WithGroupDescription(`dlv Proxy Mode:

If you invoke {{ .Command }} as "dlv" (via symlink), 
{{ .Command }} invokes dlv supporting your go and hands over stdio/stderr, args and envvar to it.

If suitable dlv is not installed, {{ .Command }} will install it.

dlvs which installed by velde are in ${VELDE_PATH} (default: ~/.velde).

{{ .Command }} mode:

If {{ .Command }} is invoked as not-"dlv", {{ .Command }} exposes itselves features.
`),
		flarc.WithSubcommand("list", cmd_list),
		flarc.WithSubcommand("uninstall", cmd_uninstall),
	)
	if err != nil {
		fmt.Fprintf(stderr, "[velde] bug. please report it: %s", err)
		return 1
	}

	return flarc.Run(
		ctx, main,
		flarc.WithInput(stdin),
		flarc.WithOutput(stdout, stderr),
		flarc.WithArgs(args),
	)
}
