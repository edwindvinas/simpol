// +build !appengine

// Package url implements url interface for anko script.
package url

import (
	u "net/url"

	"github.com/edwindvinas/simpol/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("url")
	m.Define("Parse", u.Parse)
	return m
}
