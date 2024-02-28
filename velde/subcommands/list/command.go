package list

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/youta-t/flarc"
	"github.com/youta-t/velde/velde/env"
	"github.com/youta-t/velde/velde/versions"
)

type flags struct {
	Go string `help:"list only delve compatible with go version"`
}

func Command() (flarc.Command, error) {
	return flarc.NewCommand(
		"list installed dlv by velde",
		flags{},
		flarc.Args{
			{Name: "VERSION", Help: "dlv version to be find", Repeatable: true},
		},
		func(ctx context.Context, c flarc.Commandline[flags], a []any) error {

			var gov versions.GoVersion
			checkGo := false
			if _gov := c.Flags().Go; _gov != "" {
				if !strings.HasPrefix(_gov, "go") {
					_gov = "go" + _gov
				}

				if err := (&gov).Set(_gov); err != nil {
					return err
				}
				checkGo = true
			}

			vs := make([]versions.DelveVersion, 0, len(c.Args()["VERSION"]))
			for _, v := range c.Args()["VERSION"] {
				dlvv := versions.DelveVersion{}
				err := (&dlvv).Set(v)
				if err != nil {
					return err
				}
				vs = append(vs, dlvv)
			}

			vp, err := env.VeldePath()
			if err != nil {
				return err
			}

			for _, v := range versions.Supports {
				if checkGo && !v.Compat(gov) {
					continue
				}
				if 0 < len(vs) {
					for _, vv := range vs {
						if vv.Cmp(v.Delve) != 0 {
							continue
						}
					}
				}

				fp := filepath.Join(vp, v.Delve.String())
				if _, err := os.Stat(filepath.Join(fp, "dlv")); err != nil {
					if os.IsNotExist(err) {
						continue
					}
					return err
				}

				fmt.Fprintf(c.Stdout(), "%s\t%s\n", v.Delve.String(), fp)
			}
			fmt.Fprintln(c.Stdout())

			return nil
		},
		flarc.WithDescription(`lists dlv installed by velde.

dlv are searched from ${VELDE_PATH} ("~/.velde" by default), where velde installs dlv into.
`),
	)
}
