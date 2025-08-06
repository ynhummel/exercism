package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	mu         sync.Mutex
	reader     io.Reader
	totalReads int64
	calls      int
}

type writeCounter struct {
	mu         sync.Mutex
	writer     io.Writer
	totalWrite int64
	calls      int
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriterCounter struct {
	mu         sync.Mutex
	readwriter io.ReadWriter
	totalReads int64
	totalWrite int64
	calls      int
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
	return &readWriterCounter{
		readwriter: readwriter,
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	readBytes, err := rc.reader.Read(p)
	if err != nil {
		return 0, err
	}

	rc.mu.Lock()
	rc.calls++
	rc.totalReads += int64(readBytes)
	rc.mu.Unlock()

	return readBytes, nil
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.totalReads, rc.calls
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wrote, err := wc.writer.Write(p)
	if err != nil {
		return 0, err
	}

	wc.mu.Lock()
	wc.calls++
	wc.totalWrite += int64(wrote)
	wc.mu.Unlock()

	return wrote, nil
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.totalWrite, wc.calls
}

func (rw *readWriterCounter) Read(p []byte) (int, error) {
	readBytes, err := rw.readwriter.Read(p)
	if err != nil {
		return 0, err
	}

	rw.mu.Lock()
	rw.calls++
	rw.totalReads += int64(readBytes)
	rw.mu.Unlock()

	return readBytes, nil
}

func (rw *readWriterCounter) ReadCount() (int64, int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	return rw.totalReads, rw.calls
}

func (rw *readWriterCounter) Write(p []byte) (int, error) {
	wrote, err := rw.readwriter.Write(p)
	if err != nil {
		return 0, err
	}

	rw.mu.Lock()
	rw.calls++
	rw.totalWrite += int64(wrote)
	rw.mu.Unlock()

	return wrote, nil
}

func (rw *readWriterCounter) WriteCount() (int64, int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	return rw.totalWrite, rw.calls
}
