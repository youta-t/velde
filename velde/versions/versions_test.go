package versions_test

import (
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/velde/velde/versions"
)

func TestGoVersion_Set(t *testing.T) {
	theory := func(when string, then its.Matcher[versions.GoVersion]) func(*testing.T) {
		return func(t *testing.T) {

			got := versions.GoVersion{}
			if err := (&got).Set(when); err != nil {
				t.Fatal(err)
			}
			then.Match(got).OrError(t)
		}
	}

	t.Run("go1.21.7", theory(
		"go1.21.7",
		its.EqEq(versions.GoVersion{Major: 1, Minor: 21, Release: 7}),
	))

	t.Run("go1.12", theory(
		"go1.12",
		its.EqEq(versions.GoVersion{Major: 1, Minor: 12, Release: 0}),
	))

	t.Run("go1.0.0", theory(
		"go1.0.0",
		its.EqEq(versions.GoVersion{Major: 1, Minor: 0, Release: 0}),
	))

	t.Run("go1.22rc1", theory(
		"go1.22rc2",
		its.EqEq(versions.GoVersion{Major: 1, Minor: 22, Release: 0, Suffix: "rc2"}),
	))

	t.Run("go1.19beta1", theory(
		"go1.19beta1",
		its.EqEq(versions.GoVersion{Major: 1, Minor: 19, Release: 0, Suffix: "beta1"}),
	))

	theoryErr := func(when string, then its.Matcher[error]) func(*testing.T) {
		return func(t *testing.T) {
			got := versions.GoVersion{}
			err := (&got).Set(when)
			then.Match(err).OrError(t)
		}
	}

	t.Run("go2.0.0", theoryErr("go2.0.0.", its.Error(versions.ErrParseGoVersion)))
	t.Run("go1", theoryErr("go1", its.Error(versions.ErrParseGoVersion)))
	t.Run("go1.21.1betaX", theoryErr("go1", its.Error(versions.ErrParseGoVersion)))
}

func TestGoVersion_Cmp(t *testing.T) {
	theory := func(a, b versions.GoVersion, then its.Matcher[int]) func(*testing.T) {
		return func(t *testing.T) {
			then.Match(a.Cmp(b)).OrError(t)
		}
	}

	t.Run("go1.21.0 < go1.22.0", theory(
		versions.GoVersion{Major: 1, Minor: 21, Release: 0},
		versions.GoVersion{Major: 1, Minor: 22, Release: 0},
		its.LesserThan(0),
	))

	t.Run("go1.22.0 > go1.21.0", theory(
		versions.GoVersion{Major: 1, Minor: 22, Release: 0},
		versions.GoVersion{Major: 1, Minor: 21, Release: 0},
		its.GreaterThan(0),
	))

	t.Run("go1.21.0 < go1.21.1", theory(
		versions.GoVersion{Major: 1, Minor: 21, Release: 0},
		versions.GoVersion{Major: 1, Minor: 21, Release: 1},
		its.LesserThan(0),
	))

	t.Run("go1.21.1 > go1.21.0", theory(
		versions.GoVersion{Major: 1, Minor: 21, Release: 1},
		versions.GoVersion{Major: 1, Minor: 21, Release: 0},
		its.GreaterThan(0),
	))

	t.Run("go1.21.1 == go1.21.1", theory(
		versions.GoVersion{Major: 1, Minor: 21, Release: 1},
		versions.GoVersion{Major: 1, Minor: 21, Release: 1},
		its.EqEq(0),
	))
}

func TestDelveVersion_Set(t *testing.T) {
	theory := func(when string, then its.Matcher[versions.DelveVersion]) func(*testing.T) {
		return func(t *testing.T) {
			dlvv := versions.DelveVersion{}
			if err := (&dlvv).Set(when); err != nil {
				t.Fatal(err)
			}

			then.Match(dlvv).OrError(t)
		}
	}

	t.Run("dlv@v1.22.1", theory(
		"dlv@v1.22.1",
		its.EqEq(versions.DelveVersion{Major: 1, Minor: 22, Release: 1}),
	))

	t.Run("dlv@v1.0.0-rc2", theory(
		"dlv@v1.0.0-rc2",
		its.EqEq(versions.DelveVersion{Major: 1, Minor: 0, Release: 0, Suffix: "rc2"}),
	))

	theoryErr := func(when string, then its.Matcher[error]) func(*testing.T) {
		return func(t *testing.T) {
			dlvv := versions.DelveVersion{}
			err := (&dlvv).Set(when)
			then.Match(err).OrError(t)
		}
	}

	t.Run("dlv@v2.0.0", theoryErr("dlv@v2.0.0", its.Error(versions.ErrParseDelveVersion)))
	t.Run("dlv@v1", theoryErr("dlv@v1", its.Error(versions.ErrParseDelveVersion)))
	t.Run("dlv@v1.10", theoryErr("dlv@v1.10", its.Error(versions.ErrParseDelveVersion)))
}
