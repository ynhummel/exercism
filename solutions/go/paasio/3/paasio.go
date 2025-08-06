package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	sync.Mutex
	reader     io.Reader
	totalBytes int64
	calls      int
}

type writeCounter struct {
	sync.Mutex
	writer     io.Writer
	totalBytes int64
	calls      int
}

type readWriterCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		writer: writer,
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		reader: reader,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriterCounter{NewReadCounter(readwriter), NewWriteCounter(readwriter)}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.Lock()
	defer rc.Unlock()

	readBytes, err := rc.reader.Read(p)

	rc.calls++
	rc.totalBytes += int64(readBytes)
	return readBytes, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()
	return rc.totalBytes, rc.calls
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()

	wrote, err := wc.writer.Write(p)

	wc.calls++
	wc.totalBytes += int64(wrote)
	return wrote, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.Lock()
	defer wc.Unlock()
	return wc.totalBytes, wc.calls
}
