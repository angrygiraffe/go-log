package log

import (
	"os"
)

// FixedSizeFileWriter create new log if log size exceed
type FixedSizeFileWriter struct {
	Name     string
	MaxSize  int64
	MaxCount int

	file        *os.File
	currentSize int64
}

// Write implements io.Writer
func (w *FixedSizeFileWriter) Write(p []byte) (n int, err error) {
	if w.file == nil {
		w.openNextFile()
	} else if w.currentSize > w.MaxSize {
		w.file.Close()
		w.openNextFile()
	}

	w.currentSize += int64(len(p))

	return w.file.Write(p)
}

func (w *FixedSizeFileWriter) openNextFile() (err error) {

	RotateFiles(w.Name, w.MaxCount)

	w.file, err = os.OpenFile(w.Name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	stat, err := w.file.Stat()
	if err != nil {
		return
	}
	w.currentSize = stat.Size()

	return
}
