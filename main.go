package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/youta-t/velde/velde"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(), os.Interrupt, os.Kill,
	)
	defer cancel()

	cmdname := filepath.Base(os.Args[0])
	if cmdname == "dlv" {
		// proxy mode!
		code, err := velde.Proxy(
			ctx,
			os.Stdin, os.Stdout, os.Stderr, os.Args[1:],
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[velde] runtime error: %s\n", err)
		}
		os.Exit(code)
	}

	velde.Main(
		ctx, os.Stdin, os.Stdout, os.Stderr, os.Args[1:],
	)

}
