todo comming soon

### func Run(cmd *exec.Cmd) (stdout []byte, stderr []byte, err error)

> Run sets writers for stdout and stderr, starts the specified command and
> waits for it to complete.
>
> The returned error is nil if the command runs, has no problems copying stdin,
> stdout, and stderr, and exits with a zero exit status. Otherwise, the error
> is of type Error.

### type Error error

> Error records the actual combined output of executed command, original error
> and executed cmd.

```go
type Error struct {
	// RunErr is a original occurred error.
	RunErr error
	// Cmd is a original executed command.
	Cmd *exec.Cmd
	// Output is a combined output of executing command.
	Output []byte
}
```

#### func (err *Error) Error() string

> Error returns string representation of Error type.

### func IsExitError(err error) bool

> IsExitError check that the specified error is an error about exit status.

### func GetExitStatus(err error) int

> GetExitStatus returns 0 if the specified error is not about of exit status.
> Otherwise, will be returned actual exit status.
