package cmddlv

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/youta-t/velde/velde/env"
	"github.com/youta-t/velde/velde/versions"
)

type delve struct {
	veldepath string
	version   versions.DelveVersion
	envs      []string
}

type Delve interface {
	Run(ctx context.Context, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error)
	IsInstalled() bool
	GOBIN() string
}

type Option func(*delve) *delve

func WithDelveHome(home string) Option {
	return func(d *delve) *delve {
		d.veldepath = home
		return d
	}
}

func WithEnv(envs []string) Option {
	return func(d *delve) *delve {
		newenv := make([]string, len(d.envs)+len(envs))
		copy(newenv, d.envs)
		copy(newenv[len(d.envs):], envs)
		return d
	}
}

func New(v versions.DelveVersion, option ...Option) (Delve, error) {
	vp, err := env.VeldePath()
	if err != nil {
		return nil, err
	}
	d := &delve{
		veldepath: vp,
		version:   v,
	}

	for _, o := range option {
		d = o(d)
	}
	return d, nil
}

func (d *delve) Run(ctx context.Context, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error) {
	dlv := filepath.Join(d.veldepath, d.version.String(), "dlv")

	cmd := exec.CommandContext(ctx, dlv, args...)
	defer cmd.Cancel()

	cmd.Stdin = stdin
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Env = d.envs

	err := cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}

func (d *delve) GOBIN() string {
	return filepath.Join(d.veldepath, d.version.String())
}

func (d *delve) IsInstalled() bool {
	dlv := filepath.Join(d.GOBIN(), "dlv")
	stat, err := os.Stat(dlv)
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}
