package executil

import (
	"errors"
	"reflect"
	"testing"
	"unsafe"

	"golang.org/x/crypto/ssh"

	"github.com/stretchr/testify/assert"
)

func TestGetExitStatus_ReturnsExitStatusOfActualError(
	t *testing.T,
) {
	test := assert.New(t)

	errExit1 := getCommandWithStdoutAndStderrAndExitStatus1().Run()
	test.Equal(1, GetExitStatus(errExit1))

	errExit2 := getCommandWithStdoutAndStderrAndExitStatus2().Run()
	test.Equal(2, GetExitStatus(errExit2))
}

func TestGetExitStatus_ReturnsExitStatusOfExecutilErrorWithActualError(
	t *testing.T,
) {
	test := assert.New(t)

	_, _, errExit1 := Run(getCommandWithStdoutAndStderrAndExitStatus1())
	test.Equal(1, GetExitStatus(errExit1))

	_, _, errExit2 := Run(getCommandWithStdoutAndStderrAndExitStatus2())
	test.Equal(2, GetExitStatus(errExit2))
}

func TestGetExitStatus_ReturnsZeroForNil(t *testing.T) {
	test := assert.New(t)

	test.Equal(0, GetExitStatus(nil))
}

func TestGetExitStatus_ReturnsZeroForExecutilErrorWithNil(t *testing.T) {
	test := assert.New(t)

	test.Equal(0, GetExitStatus(&Error{RunErr: nil}))
}

func TestGetExitStatus_ReturnsZeroForExecutilErrorWithNonExitError(
	t *testing.T,
) {
	test := assert.New(t)

	test.Equal(0, GetExitStatus(&Error{RunErr: errors.New("blah")}))
}

func TestGetExitStatus_ReturnsZeroForNonExitError(
	t *testing.T,
) {
	test := assert.New(t)

	test.Equal(0, GetExitStatus(errors.New("blah")))
}

func TestGetExitStatus_ReturnsExitStatusOfSshExitError(
	t *testing.T,
) {
	test := assert.New(t)

	err := new(ssh.ExitError)

	pointer := reflect.Indirect(reflect.ValueOf(err)).
		FieldByName("Waitmsg").FieldByName("status").UnsafeAddr()

	exitcode := 8
	*(*int)(unsafe.Pointer(pointer)) = *(*int)(unsafe.Pointer(&exitcode))

	test.Equal(8, GetExitStatus(err))
}
