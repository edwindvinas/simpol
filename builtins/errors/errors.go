// Package errors implements errors interface for anko script.
package errors

import (
	pkg "errors"
	"github.com/edwindvinas/simpol/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewModule("errors")
	m.Define("New", pkg.New)
	return m
}
