package controller

import (
	"github.com/mosen/openshift-janitor-operator/pkg/controller/janitor"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, janitor.Add)
}
