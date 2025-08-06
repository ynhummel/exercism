package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	sync.Mutex
	reader     io.Reader
	totalReads int64
	calls      int
}

type writeCounter struct {
	sync.Mutex
	writer     io.Writer
	totalWrite int64
	calls      int
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriterCounter struct {
	sync.Mutex
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
	rc.Lock()
	defer rc.Unlock()

	readBytes, err := rc.reader.Read(p)
	if err != nil {
		return 0, err
	}

	rc.calls++
	rc.totalReads += int64(readBytes)
	return readBytes, nil
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()
	return rc.totalReads, rc.calls
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()

	wrote, err := wc.writer.Write(p)
	if err != nil {
		return 0, err
	}

	wc.calls++
	wc.totalWrite += int64(wrote)
	return wrote, nil
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.Lock()
	defer wc.Unlock()
	return wc.totalWrite, wc.calls
}

func (rw *readWriterCounter) Read(p []byte) (int, error) {
	rw.Lock()
	defer rw.Unlock()

	readBytes, err := rw.readwriter.Read(p)
	if err != nil {
		return 0, err
	}

	rw.calls++
	rw.totalReads += int64(readBytes)
	return readBytes, nil
}

func (rw *readWriterCounter) ReadCount() (int64, int) {
	rw.Lock()
	defer rw.Unlock()
	return rw.totalReads, rw.calls
}

func (rw *readWriterCounter) Write(p []byte) (int, error) {
	rw.Lock()
	defer rw.Unlock()

	wrote, err := rw.readwriter.Write(p)
	if err != nil {
		return 0, err
	}

	rw.calls++
	rw.totalWrite += int64(wrote)
	return wrote, nil
}

func (rw *readWriterCounter) WriteCount() (int64, int) {
	rw.Lock()
	defer rw.Unlock()
	return rw.totalWrite, rw.calls
}
