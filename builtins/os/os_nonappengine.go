// +build !appengine

package os

import (
	"github.com/edwindvinas/simpol/vm"
	pkg "os"
	"reflect"
)

func handleAppEngine(m *vm.Env) {
	m.Define("Getppid", reflect.ValueOf(pkg.Getppid))
}
