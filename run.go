package executil

import (
	"bytes"
	"io"
	"os/exec"
)

type option int

const (
	// IgnoreStdout is option for executil.Run() which should be passed if
	// stdout data should be ignored.
	IgnoreStdout option = iota

	// IgnoreStderr is option for executil.Run() which should be passed if
	// stderr data should be ignored.
	IgnoreStderr
)

// Run sets writers for stdout and stderr, starts the specified command and
// waits for it to complete.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
// Otherwise, the error is of type Error.
func Run(
	cmd *exec.Cmd, options ...option,
) (stdout []byte, stderr []byte, err error) {
	var (
		stdoutBuffer   = &bytes.Buffer{}
		stderrBuffer   = &bytes.Buffer{}
		combinedBuffer = &threadsafeBuffer{}

		ignoreStdout bool
		ignoreStderr bool
	)

	for _, option := range options {
		switch option {
		case IgnoreStdout:
			ignoreStdout = true

		case IgnoreStderr:
			ignoreStderr = true
		}
	}

	if !ignoreStdout {
		cmd.Stdout = io.MultiWriter(stdoutBuffer, combinedBuffer)
	}
	if !ignoreStderr {
		cmd.Stderr = io.MultiWriter(stderrBuffer, combinedBuffer)
	}

	runErr := cmd.Run()
	if runErr != nil {
		err = &Error{
			RunErr: runErr,
			Cmd:    cmd,
			Output: combinedBuffer.Bytes(),
		}
	}

	return stdoutBuffer.Bytes(), stderrBuffer.Bytes(), err
}
