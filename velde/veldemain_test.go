package velde_test

import (
	"bytes"
	"context"
	"strings"

	"github.com/youta-t/velde/velde"
)

func ExampleMain() {
	ctx := context.Background()
	in := new(bytes.Buffer)
	out := new(strings.Builder)
	err := new(strings.Builder)
	velde.Main(ctx, in, out, err, []string{"-h"})
	// Output:
}
