package cmdgo

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/youta-t/velde/velde/versions"
)

type cmdgo struct {
	goroot string
	name   string
	env    []string
}

type Go interface {
	Version(context.Context) (versions.GoVersion, int, error)
	Install(ctx context.Context, stdout io.Writer, stderr io.Writer, packageName string) (int, error)
}

type Option func(*cmdgo) *cmdgo

func WithGoroot(grt string) Option {
	return func(c *cmdgo) *cmdgo {
		c.goroot = grt
		return c
	}
}

func WithEnv(envs []string) Option {
	return func(c *cmdgo) *cmdgo {
		newenvs := make([]string, len(c.env)+len(envs))
		copy(newenvs, c.env)
		copy(newenvs[len(c.env):], envs)
		c.env = newenvs
		return c
	}
}

func New(option ...Option) Go {
	c := &cmdgo{
		goroot: os.Getenv("GOROOT"),
	}

	for _, o := range option {
		c = o(c)
	}

	if c.goroot != "" {
		c.name = filepath.Join(c.goroot, "bin", "go")
	} else {
		c.name = "go" // find from PATH
	}

	return c
}

func (g *cmdgo) run(
	ctx context.Context,
	stdout io.Writer,
	stderr io.Writer,
	args ...string,
) (int, error) {
	cmd := exec.CommandContext(ctx, "go", args...)
	defer cmd.Cancel()

	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = append([]string{}, os.Environ()...)
	cmd.Env = append(cmd.Env, g.env...)

	err := cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}

func (g *cmdgo) Version(ctx context.Context) (versions.GoVersion, int, error) {
	stdout := new(strings.Builder)
	code, err := g.run(ctx, stdout, io.Discard, "env", "GOVERSION")
	if err != nil {
		return versions.GoVersion{}, code, err
	}

	gov := versions.GoVersion{}
	err = (&gov).Set(strings.Trim(stdout.String(), "\n"))
	return gov, code, err
}

func (g *cmdgo) Install(
	ctx context.Context,
	stdout io.Writer,
	stderr io.Writer,
	packageName string,
) (int, error) {
	if stdout == nil {
		stdout = io.Discard
	}
	if stderr == nil {
		stderr = io.Discard
	}
	return g.run(ctx, stdout, stderr, "install", packageName)
}
