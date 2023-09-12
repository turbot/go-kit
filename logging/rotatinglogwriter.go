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

type RotatingLogWriter struct {
	directory string
	prefix    string

	currentWriter io.Writer
	currentPath   string

	rotateLock sync.Mutex
}

func NewRotatingLogWriter(directory string, prefix string) *RotatingLogWriter {
	return &RotatingLogWriter{
		directory: directory,
		prefix:    prefix,
	}
}
func (dwr *RotatingLogWriter) rotateLogTarget(targetPath string) (err error) {
	dwr.rotateLock.Lock()
	defer dwr.rotateLock.Unlock()

	// update to the current path
	dwr.currentPath = targetPath

	// check if the file actually doesn't exist
	// another thread may have created it while we were waiting for the lock
	if !files.FileExists(targetPath) {
		// we need to flush the current one
		// try to cast it to a Closer (if this is nil, isCloseable will be false)
		closeableWriter, isCloseable := dwr.currentWriter.(io.Closer)
		if isCloseable {
			closeableWriter.Close()
		}
	}

	// we could be in here because the file exists,
	// but we are starting up for the first time
	if dwr.currentWriter == nil {
		// create a new one
		dwr.currentWriter, err = os.OpenFile(dwr.currentPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open steampipe log file: %s", err.Error())
		}
	}
	return nil
}

func (dwr *RotatingLogWriter) Write(p []byte) (n int, err error) {
	pathShouldBe := filepath.Join(dwr.directory, fmt.Sprintf("%s-%s.log", dwr.prefix, time.Now().Format(time.DateOnly)))

	// the code outside of this IF block should be simple and blazing fast
	//
	// the code inside the IF block will probably execute once in 24 hours at most
	// for an instance, but the code outside is used by every log line
	if dwr.currentPath != pathShouldBe {
		if err := dwr.rotateLogTarget(pathShouldBe); err != nil {
			return 0, err
		}
	}

	return dwr.currentWriter.Write(p)
}
