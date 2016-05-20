package executil

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"golang.org/x/crypto/ssh"

	"github.com/stretchr/testify/assert"
)

func TestIsExitError_ReturnsTrueForOsExecExitError(t *testing.T) {
	test := assert.New(t)

	test.True(
		IsExitError(
			&exec.ExitError{ProcessState: &os.ProcessState{}},
		),
	)
}

func TestIsExitError_ReturnsTrueForExecutilErrorWhenRunIsOsExecExitError(
	t *testing.T,
) {
	test := assert.New(t)

	test.True(
		IsExitError(
			&Error{RunErr: &exec.ExitError{ProcessState: &os.ProcessState{}}},
		),
	)
}

func TestIsExitError_ReturnsTrueForSshExitError(t *testing.T) {
	test := assert.New(t)

	test.True(
		IsExitError(new(ssh.ExitError)),
	)
}

func TestIsExitError_ReturnsFalseForExecutilErrorWithNonExecError(
	t *testing.T,
) {
	test := assert.New(t)

	test.False(IsExitError(&Error{RunErr: errors.New("blaaa")}))
}

func TestIsExitError_ReturnsFalseForExecutilErrorWithNil(
	t *testing.T,
) {
	test := assert.New(t)

	test.False(IsExitError(&Error{RunErr: nil}))
}

func TestIsExitError_ReturnsFalseForAllErrors(
	t *testing.T,
) {
	test := assert.New(t)

	test.False(IsExitError(errors.New("u shall not pass")))
}

func TestIsExitError_ReturnsFalseForNil(
	t *testing.T,
) {
	test := assert.New(t)

	test.False(IsExitError(nil))
}
