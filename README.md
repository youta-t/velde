velde -- delve version switcher
===============================

For time-traveling gophers.

Motivation
----------

By `go install`, we can have just only single `dlv`.
However, `dlv` of a version can supports some range of go versions, not all go.
If you are writing old and new go, managing `dlv` can be toil.

velde, this command, manages many versions of `dlv`, install them as needed, and proxies to suitable one for your go!

Install
------------

First, uninstall `dlv`.

Second, `go install`.

```
go install github.com/youta-t/velde
```

Finally, make symlink to velde as `dlv`.

```
ln -s ./velde ~/go/bin/dlv
```

Use
----

velde has 2 modes.

### `dlv` Proxy Mode

When velde is invoked as `dlv` (via symlink), it perform 2 tasks in order.

1. installs `dlv` which supports `go` you using.
2. hands over tasks to the `dlv` supports your `go`.

Task 1 will be skipped if you have suitable `dlv` for your `go`. 

Installed `dlv` by velde are in `${VELDE_PATH}` (default; `~/.velde`).

#### go sdk detection

`go` is searched from `${PATH}`.

velde invokes `go version` as a subprocess to know the version of go which is used,
and involes `go install` as a subprocess (with overwriting `GOBIN`) to install `dlv`.

### velde mode

When velde is invoked as any name but `dlv`, it perform management functions for your `dlv`.

There are several subcommands.

- `velde list`: list installed dlv by velde
- `velde uninstall`: uninstall delve installed by velde
    - Also, you can delete `${VELDE_PATH}/dlv@${VERSION}`. It does this.

For each subcommand, `dlv`'s version are in format `dlv@v${MAJOR}.${MINOR}.${RELEASE}[-${SUFFIX}]`.
