package controller

import (
	"github.com/mosen/openshift-janitor-operator/pkg/controller/sweep"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sweep.Add)
}
