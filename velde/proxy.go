package velde

import (
	"context"
	"fmt"
	"io"

	"github.com/youta-t/velde/velde/cmddlv"
	"github.com/youta-t/velde/velde/cmdgo"
	"github.com/youta-t/velde/velde/versions"
)

func Proxy(
	ctx context.Context,

	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,

	args []string,
) (int, error) {
	g := cmdgo.New()

	gover, code, err := g.Version(ctx)
	if err != nil || code != 0 {
		return code, err
	}

	var dlvv versions.DelveVersion
	found := false
	for _, sup := range versions.Supports {
		if sup.Compat(gover) {
			dlvv = sup.Delve
			found = true
			break
		}
	}
	if !found {
		return 1, fmt.Errorf("[velde] %s seems not supported", gover)
	}

	dlv, err := cmddlv.New(dlvv)
	if err != nil {
		return 1, err
	}

	if !dlv.IsInstalled() {
		fmt.Fprintf(stderr, "[velde] dlv supporting %s is not found.\n", gover)
		fmt.Fprintf(stderr, "[velde] Install %s into %s...\n", dlvv, dlv.GOBIN())
		{
			g := cmdgo.New(cmdgo.WithEnv([]string{
				fmt.Sprintf("GOBIN=%s", dlv.GOBIN()),
			}))
			code, err := g.Install(ctx, stdout, stderr, dlvv.PackageQual())
			if code != 0 || err != nil {
				return code, err
			}
		}
		fmt.Fprintf(stderr, "[velde] %s is installed at %s\n", dlvv, dlv.GOBIN())
		fmt.Fprintln(stderr, "[velde] starting dlv...")
	}

	return dlv.Run(ctx, stdin, stdout, stderr, args...)
}
