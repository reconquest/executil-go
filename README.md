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


--------------------------------------------------------------------------

#### coverage: 100.0% of statements

```
github.com/kovetskiy/executil/error.go:20:      Error           100.0%
github.com/kovetskiy/executil/exit_error.go:8:  getWaitStatus   100.0%
github.com/kovetskiy/executil/exit_error.go:27: IsExitError     100.0%
github.com/kovetskiy/executil/exit_error.go:33: GetExitStatus   100.0%
github.com/kovetskiy/executil/run.go:15:        Write           100.0%
github.com/kovetskiy/executil/run.go:29:        Run             100.0%
total:                                          (statements)    100.0%
````
| Subject       | Behavior                          |
| ------------- | --------------------------------- |
| Run           | Returns Stdout If Stderr Is Empty |
| Run           | Returns Stderr If Stdout Is Empty |
| Run           | Returns Stdout And Stderr If Both Not Empty |
| Run           | Returns Error If Command Failed |
| Run           | Returns Error Type Of Executil Error |
| Run           | Returns Error With Exit Status1 |
| Run           | Returns Error With Exit Status2 |
| Error.Output  | Is Combined Output Of Executed Command |
| Error.Cmd     | Is Actual Executed Command |
| Error.RunErr  | Is Actual Error Of Executed Command |
| Error->Error  | Contains Of Actual Error |
| Error->Error  | Contains Of Command Args |
| Error->Error  | Contains Of Output If Output Is Not Empty |
| Error->Error  | Contains Of Message With Output If Output Is Not Empty |
| Error->Error  | Contains Of Message Without Output If Output Is Empty |
| IsExitError   | Returns True For Os Exit Error |
| IsExitError   | Returns True For Executil Error When Run Err Is Exit Error |
| IsExitError   | Returns False For Executil Error With Non Exec Error |
| IsExitError   | Returns False For Executil Error With Nil |
| IsExitError   | Returns False For All Errors |
| IsExitError   | Returns False For Nil |
| GetExitStatus | Returns Exit Status Of Actual Error |
| GetExitStatus | Returns Exit Status Of Executil Error With Actual Error |
| GetExitStatus | Returns Zero For Nil |
| GetExitStatus | Returns Zero For Executil Error With Nil |
| GetExitStatus | Returns Zero For Executil Error With Non Exit Error |
| GetExitStatus | Returns Zero For Non Exit Error |
