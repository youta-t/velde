//go:generate go run github.com/youta-t/its/structer
package versions

import (
	"cmp"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type GoVersion struct {
	Major   int
	Minor   int
	Release int
	Suffix  string
}

func Go(major, minor, release int, suffix string) GoVersion {
	return GoVersion{
		Major:   major,
		Minor:   minor,
		Release: release,
		Suffix:  suffix,
	}
}

var ErrPerseVersion = errors.New("unrecognizeable version")
var ErrParseGoVersion = fmt.Errorf("%w: go version", ErrPerseVersion)
var ErrParseDelveVersion = fmt.Errorf("%w: delve version", ErrPerseVersion)

func (gv *GoVersion) Set(value string) error {
	if !strings.HasPrefix(value, "go1.") {
		return fmt.Errorf("%w: %s", ErrParseGoVersion, value)
	}
	gv.Major = 1
	minor, release, ok := strings.Cut(value[4:], ".")
	if ok {
		if iminor, err := strconv.Atoi(minor); err != nil {
			return fmt.Errorf("%w: %s", ErrParseGoVersion, value)
		} else {
			gv.Minor = iminor
		}
		if ir, err := strconv.Atoi(release); err != nil {
			return fmt.Errorf("%w: %s", ErrParseGoVersion, value)
		} else {
			gv.Release = ir
		}

		return nil
	}

	gv.Release = 0

	gv.Suffix = strings.TrimLeft(minor, "0123456789")
	minor = minor[:len(minor)-len(gv.Suffix)]
	if iminor, err := strconv.Atoi(minor); err != nil {
		return fmt.Errorf("%w: %s", ErrParseGoVersion, value)
	} else {
		gv.Minor = iminor
	}

	return nil
}

func (gv GoVersion) String() string {
	if gv.Suffix != "" {
		return fmt.Sprintf("go%d.%d%s", gv.Major, gv.Minor, gv.Suffix)
	}
	return fmt.Sprintf("go%d.%d.%d", gv.Major, gv.Minor, gv.Release)
}

func (gv GoVersion) Cmp(other GoVersion) int {
	switch c := cmp.Compare(gv.Major, other.Major); c {
	case 0:
	default:
		return c
	}
	switch c := cmp.Compare(gv.Minor, other.Minor); c {
	case 0:
	default:
		return c
	}
	switch c := cmp.Compare(gv.Release, other.Release); c {
	case 0:
	default:
		return c
	}

	gvbeta := strings.HasSuffix(gv.Suffix, "beta")
	otherbeta := strings.HasSuffix(other.Suffix, "beta")
	gvrc := strings.HasSuffix(gv.Suffix, "rc")
	otherrc := strings.HasSuffix(other.Suffix, "rc")

	if otherrc && gvbeta {
		return -1
	}
	if gvrc && otherbeta {
		return 1
	}

	return cmp.Compare(gv.Suffix, other.Suffix)
}

func Delve(major, minor, release int, suffix string) DelveVersion {
	return DelveVersion{
		Major:   major,
		Minor:   minor,
		Release: release,
		Suffix:  suffix,
	}
}

type DelveVersion struct {
	Major   int
	Minor   int
	Release int
	Suffix  string
}

func (dlvv *DelveVersion) Set(value string) error {
	if !strings.HasPrefix(value, "dlv@v1.") {
		return fmt.Errorf("%w: %s", ErrParseDelveVersion, value)
	}
	dlvv.Major = 1
	minor, release, ok := strings.Cut(value[7:], ".")
	if !ok {
		return fmt.Errorf("%w: %s", ErrParseDelveVersion, value)
	}

	if iminor, err := strconv.Atoi(minor); err != nil {
		return fmt.Errorf("%w: delve version: %s", ErrParseDelveVersion, value)
	} else {
		dlvv.Minor = iminor
	}

	release, dlvv.Suffix, _ = strings.Cut(release, "-")
	if ir, err := strconv.Atoi(release); err != nil {
		return fmt.Errorf("%w: %s", ErrParseDelveVersion, value)
	} else {
		dlvv.Release = ir
	}

	return nil
}

func (dlvv DelveVersion) String() string {
	if dlvv.Suffix != "" {
		return fmt.Sprintf("dlv@v%d.%d.%d-%s", dlvv.Major, dlvv.Minor, dlvv.Release, dlvv.Suffix)
	}
	return fmt.Sprintf("dlv@v%d.%d.%d", dlvv.Major, dlvv.Minor, dlvv.Release)
}

func (dlvv DelveVersion) PackageQual() string {
	return "github.com/go-delve/delve/cmd/" + dlvv.String()
}

func (dlvv DelveVersion) Cmp(other DelveVersion) int {
	switch c := cmp.Compare(dlvv.Major, other.Major); c {
	case 0:
	default:
		return c
	}
	switch c := cmp.Compare(dlvv.Minor, other.Minor); c {
	case 0:
	default:
		return c
	}
	switch c := cmp.Compare(dlvv.Release, other.Release); c {
	case 0:
	default:
		return c
	}

	gvbeta := strings.HasSuffix(dlvv.Suffix, "alpha")
	otherbeta := strings.HasSuffix(other.Suffix, "alpha")
	// delve have not released as beta yet.
	gvrc := strings.HasSuffix(dlvv.Suffix, "rc")
	otherrc := strings.HasSuffix(other.Suffix, "rc")

	if otherrc && gvbeta {
		return -1
	}
	if gvrc && otherbeta {
		return 1
	}

	return cmp.Compare(dlvv.Suffix, other.Suffix)
}
