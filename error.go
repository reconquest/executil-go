package executil

import (
	"fmt"
	"os/exec"

	"github.com/seletskiy/hierr"
)

// Error records the actual combined output of executed command, original error
// and executed cmd.
type Error struct {
	// RunErr is a original occurred error.
	RunErr error
	// Cmd is a original executed command.
	Cmd *exec.Cmd
	// Output is a combined output of executing command.
	Output []byte
}

// Error returns string representation of Error type.
func (err *Error) Error() string {
	value := fmt.Sprintf("exec %q error (%s) ", err.Cmd.Args, err.RunErr)
	if len(err.Output) > 0 {
		value = value + "with output:\n" + string(err.Output)
	}
	return value
}

// HierarchicalError returns hierarchical string representation using hierr
// package.
func (err *Error) HierarchicalError() string {
	runError := err.RunErr
	if len(err.Output) > 0 {
		runError = hierr.Push(runError, string(err.Output))
	}

	return hierr.Errorf(runError, "exec %q error", err.Cmd.Args).Error()
}
