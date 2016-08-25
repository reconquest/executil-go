package executil

import (
	"fmt"
	"os/exec"

	"github.com/reconquest/ser-go"
)

type CommandWithArgs interface {
	GetArgs() []string
}

// Error records the actual combined output of executed command, original error
// and executed cmd.
type Error struct {
	// RunErr is a original occurred error.
	RunErr error

	// Cmd is a original executed command.
	// can be *exec.Command or lexec.Command
	Cmd interface{}

	// Output is a combined output of executing command.
	Output []byte
}

// Error returns string representation of Error type.
func (err *Error) Error() string {
	args := err.getArgs()
	if len(args) == 0 {
		return err.RunErr.Error()
	}

	value := fmt.Sprintf("exec %q error (%s) ", args, err.RunErr)
	if len(err.Output) > 0 {
		value = value + "with output:\n" + string(err.Output)
	}
	return value
}

// HierarchicalError returns hierarchical string representation using hierr
// package.
func (err *Error) HierarchicalError() string {
	args := err.getArgs()
	if len(args) == 0 {
		return err.RunErr.Error()
	}

	runError := err.RunErr
	if len(err.Output) > 0 {
		runError = ser.Push(runError, string(err.Output))
	}

	return ser.Errorf(runError, "exec %q error", args).Error()
}

func (err *Error) getArgs() []string {
	if cmd, ok := err.Cmd.(CommandWithArgs); ok {
		return cmd.GetArgs()
	} else if cmd, ok := err.Cmd.(*exec.Cmd); ok {
		return cmd.Args
	}

	return []string{}
}
