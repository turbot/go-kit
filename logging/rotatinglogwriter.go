package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/turbot/go-kit/files"
)

// RotatingLogWriter is an io.Writer which can rotate the log files that it is writing to
// Target files are rotated once per day according to the local system time
type RotatingLogWriter struct {
	directory string
	prefix    string

	currentWriter io.Writer
	currentPath   string

	rotateLock sync.Mutex
}

// RotatingLogWriter returns io.Writer which always writes to a file in the given `directory`.
// The file is named `{prefix}-YYYY-MM-DD.log` and the log files are rotated once per day as per the system time
func NewRotatingLogWriter(directory string, prefix string) *RotatingLogWriter {
	return &RotatingLogWriter{
		directory: directory,
		prefix:    prefix,
	}
}

func (w *RotatingLogWriter) Write(p []byte) (n int, err error) {
	expectedPath := filepath.Join(w.directory, fmt.Sprintf("%s-%s.log", w.prefix, time.Now().Format(time.DateOnly)))

	// the code outside of this IF block should be simple and blazing fast
	//
	// the code inside the IF block will probably execute once in 24 hours at most
	// for an instance, but the code outside is used by every log line
	if w.currentPath != expectedPath {
		if err := w.rotateLogTarget(expectedPath); err != nil {
			return 0, err
		}
		// update the current path
		w.currentPath = expectedPath

		// update the writer
		w.currentWriter, err = os.OpenFile(w.currentPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			err = fmt.Errorf("failed to open steampipe log file: %s", err.Error())
			return 0, err
		}
	}

	return w.currentWriter.Write(p)
}

func (w *RotatingLogWriter) rotateLogTarget(targetPath string) (err error) {
	w.rotateLock.Lock()
	defer w.rotateLock.Unlock()

	// check if the file actually doesn't exist
	if files.FileExists(targetPath) {
		// nothing to do here
		// another thread may have created it while we were waiting for the lock
		return nil
	}

	// we need to flush the current one
	// try to cast it to a Closer (if this is nil, isCloseable will be false)
	closeableWriter, isCloseable := w.currentWriter.(io.Closer)
	if isCloseable {
		closeableWriter.Close()
	}
	return nil
}
