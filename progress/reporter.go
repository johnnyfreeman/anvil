package progress

import "errors"

// ProgressReporter defines the interface for reporting task progress.
type ProgressReporter interface {
	Start(task string)
	Success(task string)
	Failure(task string, err error)
	Done()
}

var ErrCancelled = errors.New("operation cancelled")
