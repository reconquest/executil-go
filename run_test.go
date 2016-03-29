package executil

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun_ReturnsStdoutIfStderrIsEmpty(t *testing.T) {
	test := assert.New(t)

	stdout, stderr, err := Run(getCommandWithStdout())
	test.NoError(err)
	test.Equal("stdout\n", string(stdout))
	test.Equal("", string(stderr))
}

func TestRun_ReturnsStderrIfStdoutIsEmpty(t *testing.T) {
	test := assert.New(t)

	stdout, stderr, err := Run(getCommandWithStderr())
	test.NoError(err)
	test.Equal("", string(stdout))
	test.Equal("stderr\n", string(stderr))
}

func TestRun_ReturnsStdoutAndStderrIfBothNotEmpty(t *testing.T) {
	test := assert.New(t)

	stdout, stderr, err := Run(getCommandWithStdoutAndStderr())
	test.NoError(err)
	test.Equal("stdout\n", string(stdout))
	test.Equal("stderr\n", string(stderr))
}

func TestRun_ReturnsErrorIfCommandFailed(t *testing.T) {
	test := assert.New(t)

	_, _, err := Run(getCommandWithUnknownBinary())
	test.Error(err)
}

func TestRun_ReturnsErrorTypeOfExecutilError(t *testing.T) {
	test := assert.New(t)

	_, _, err := Run(getCommandWithUnknownBinary())
	test.IsType(&Error{}, err)
}

func TestRun_ReturnsErrorWithExitStatus1(t *testing.T) {
	test := assert.New(t)

	stdout, stderr, err := Run(getCommandWithStdoutAndStderrAndExitStatus1())
	test.Error(err)
	test.Equal("stdout\n", string(stdout))
	test.Equal("stderr\n", string(stderr))
	test.IsType(&exec.ExitError{}, err.(*Error).RunErr)
	test.True(IsExitError(err), "error should be type of ExitError")
	test.Equal(GetExitStatus(err), 1)
}

func TestRun_ReturnsErrorWithExitStatus2(t *testing.T) {
	test := assert.New(t)

	stdout, stderr, err := Run(getCommandWithStdoutAndStderrAndExitStatus2())
	test.Error(err)
	test.Equal("stdout\n", string(stdout))
	test.Equal("stderr\n", string(stderr))
	test.IsType(&Error{}, err)
	test.IsType(&exec.ExitError{}, err.(*Error).RunErr)
	test.True(IsExitError(err), "error should be type of ExitError")
	test.Equal(GetExitStatus(err), 2)
}
