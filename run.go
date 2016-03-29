package executil

import (
	"bytes"
	"io"
	"os/exec"
	"sync"
)

type threadsafeBuffer struct {
	sync.Mutex
	bytes.Buffer
}

func (buffer *threadsafeBuffer) Write(data []byte) (int, error) {
	buffer.Lock()
	defer buffer.Unlock()

	return buffer.Buffer.Write(data)
}

// Run sets writers for stdout and stderr, starts the specified command and
// waits for it to complete.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
// Otherwise, the error is of type Error.
func Run(cmd *exec.Cmd) (stdout []byte, stderr []byte, err error) {
	var (
		stdoutBuffer   = &bytes.Buffer{}
		stderrBuffer   = &bytes.Buffer{}
		combinedBuffer = &threadsafeBuffer{}
	)

	cmd.Stdout = io.MultiWriter(stdoutBuffer, combinedBuffer)
	cmd.Stderr = io.MultiWriter(stderrBuffer, combinedBuffer)

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
