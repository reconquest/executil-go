package executil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Output_IsCombinedOutputOfExecutedCommand(t *testing.T) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		if !command.isExitError {
			continue
		}

		_, _, err := Run(command.get())
		test.IsType(&Error{}, err)
		test.Equal("stdout\nstderr\n", string(err.(*Error).Output))
	}
}

func TestError_Cmd_IsActualExecutedCommand(t *testing.T) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		cmd := command.get()
		_, _, err := Run(cmd)
		test.IsType(&Error{}, err)
		test.Equal(cmd, err.(*Error).Cmd)
	}
}

func TestError_RunErr_IsActualErrorOfExecutedCommand(t *testing.T) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		_, _, err := Run(command.get())
		test.IsType(&Error{}, err)

		test.Equal(
			command.get().Run().Error(), err.(*Error).RunErr.Error(),
		)
	}
}

func TestError_Error_ContainsOfActualError(t *testing.T) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		_, _, err := Run(command.get())
		test.Contains(err.Error(), command.get().Run().Error())
	}
}

func TestError_Error_ContainsOfCommandArgs(t *testing.T) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		_, _, err := Run(command.get())
		test.Contains(err.Error(), fmt.Sprintf("%q", command.get().Args))
	}
}

func TestError_Error_ContainsOfOutputIfOutputIsNotEmpty(t *testing.T) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		if !command.isExitError {
			continue
		}

		_, _, err := Run(command.get())
		expectedOutput, _ := command.get().CombinedOutput()
		test.Contains(err.Error(), string(expectedOutput))
	}
}

func TestError_Error_ContainsOfMessageWithOutputIfOutputIsNotEmpty(
	t *testing.T,
) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		if !command.isExitError {
			continue
		}

		_, _, err := Run(command.get())
		test.Contains(err.Error(), "with output")
		test.NotContains(err.Error(), "without output")
	}
}

func TestError_Error_ContainsOfMessageWithoutOutputIfOutputIsEmpty(
	t *testing.T,
) {
	test := assert.New(t)

	for _, command := range commandsWithError {
		if command.isExitError {
			continue
		}

		_, _, err := Run(command.get())
		test.Contains(err.Error(), "without output")
		test.NotContains(err.Error(), "with output")
	}
}
