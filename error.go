package executil

import (
	"fmt"
	"os/exec"
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
	value := fmt.Sprintf("exec [%q] failed (%s) ", err.Cmd.Args, err.RunErr)
	if len(err.Output) > 0 {
		value = value + "with output:\n" + string(err.Output)
	} else {
		value = value + "without output"
	}
	return value
}
