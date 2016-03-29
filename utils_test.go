package executil

import "os/exec"

var commandsWithError = []struct {
	get         func() *exec.Cmd
	isExitError bool
}{
	{getCommandWithUnknownBinary, false},
	{getCommandWithStdoutAndStderrAndExitStatus1, true},
	{getCommandWithStdoutAndStderrAndExitStatus2, true},
}

func getCommandWithStdoutAndStderr() *exec.Cmd {
	return exec.Command(
		"bash", "-c", "echo stdout; sleep 0.05;  echo stderr >&2",
	)
}

func getCommandWithStdout() *exec.Cmd {
	return exec.Command(
		"bash", "-c", "echo stdout;",
	)
}

func getCommandWithStderr() *exec.Cmd {
	return exec.Command(
		"bash", "-c", "echo stderr >&2",
	)
}

func getCommandWithStdoutAndStderrAndExitStatus1() *exec.Cmd {
	return exec.Command(
		"bash", "-c", "echo stdout; sleep 0.05; echo stderr >&2; exit 1",
	)
}

func getCommandWithStdoutAndStderrAndExitStatus2() *exec.Cmd {
	return exec.Command(
		"bash", "-c", "echo stdout; sleep 0.05; echo stderr >&2; exit 2",
	)
}

func getCommandWithUnknownBinary() *exec.Cmd {
	return exec.Command(
		"___i_am_realy_does_not_exist_on_the_system", "--flag",
	)
}
