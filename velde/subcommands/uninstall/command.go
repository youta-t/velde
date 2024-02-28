package uninstall

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/youta-t/flarc"
	"github.com/youta-t/velde/velde/env"
	"github.com/youta-t/velde/velde/versions"
)

func Command() (flarc.Command, error) {
	return flarc.NewCommand[struct{}](
		"uninstall delve installed by velde",
		struct{}{},
		flarc.Args{
			{Name: "DELVE_VERSION", Help: "delve version to be uninstalled", Repeatable: true},
		},
		func(ctx context.Context, c flarc.Commandline[struct{}], a []any) error {
			dlvvs := []versions.DelveVersion{}
			for _, v := range c.Args()["DELVE_VERSION"] {
				dlvv := versions.DelveVersion{}
				if err := (&dlvv).Set(v); err != nil {
					fmt.Fprintf(c.Stderr(), "[velde] unknwon delve: %s\n", v)
					return err
				}
				idx, found := slices.BinarySearchFunc(
					versions.Supports, dlvv,
					func(e versions.Support, dv versions.DelveVersion) int { return -e.Delve.Cmp(dv) },
				)

				if !found {
					fmt.Fprintf(c.Stderr(), "[velde] unknwon delve: %s\n", dlvv)
					continue
				}

				dlvvs = append(dlvvs, versions.Supports[idx].Delve)
			}

			vp, err := env.VeldePath()
			if err != nil {
				return nil
			}

			for _, dlvv := range dlvvs {
				target := filepath.Join(vp, dlvv.String())
				fmt.Fprintf(c.Stderr(), "[delve] removing %s (%s)\n", dlvv, target)
				if err := os.RemoveAll(target); err != nil {
					return err
				}
			}

			return nil
		},
	)
}
