package executil

import (
	"os/exec"
	"syscall"
)

func getWaitStatus(err error) *syscall.WaitStatus {
	if err == nil {
		return nil
	}

	if utilErr, ok := err.(*Error); ok {
		err = utilErr.RunErr
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			return &waitStatus
		}
	}

	return nil
}

// IsExitError check that the specified error is an error about exit status.
func IsExitError(err error) bool {
	return getWaitStatus(err) != nil
}

// GetExitStatus returns 0 if the specified error is not about of exit status.
// Otherwise, will be returned actual exit status.
func GetExitStatus(err error) int {
	waitStatus := getWaitStatus(err)
	if waitStatus == nil {
		return 0
	}

	return waitStatus.ExitStatus()
}
