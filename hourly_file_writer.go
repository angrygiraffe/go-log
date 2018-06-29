package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// HourlyFileWriter create new log for every day
type HourlyFileWriter struct {
	Name     string
	MaxCount int
	MaxSize  int64

	file         *os.File
	nextHourTime int64
	currentSize  int64
}

// Write implements io.Writer
func (w *HourlyFileWriter) Write(p []byte) (n int, err error) {
	now := time.Now()

	if w.file == nil {
		w.openFile(&now)
	} else if now.Unix() >= w.nextHourTime || w.currentSize > w.MaxSize {
		w.file.Close()
		w.openFile(&now)
	}

	w.currentSize += int64(len(p))
	return w.file.Write(p)
}

func (w *HourlyFileWriter) openFile(now *time.Time) (err error) {
	name := fmt.Sprintf("%s.%s", w.Name, now.Format("2006010215"))

	// remove symbol link if exist
	os.Remove(w.Name)

	// rotate file
	RotateFiles(name, w.MaxCount)

	// create file
	w.file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	stat, err := w.file.Stat()
	if err != nil {
		return
	}

	w.currentSize = stat.Size()

	// create symbol
	err = os.Symlink(path.Base(name), w.Name)
	if err != nil {
		return err
	}

	nextHourTime := now.Add(time.Hour)
	year, month, day := nextHourTime.Date()
	hour := nextHourTime.Hour()
	w.nextHourTime = time.Date(year, month, day, hour, 0, 0, 0, now.Location()).Unix()

	return nil
}

// clean old files
func (w *HourlyFileWriter) cleanFiles() {
	dir := path.Dir(w.Name)

	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	prefix := path.Base(w.Name) + "."

	var matches []string
	for _, f := range fileList {
		if !f.IsDir() && strings.HasPrefix(f.Name(), prefix) {
			matches = append(matches, f.Name())
		}
	}

	if len(matches) > w.MaxCount {
		sort.Sort(sort.Reverse(sort.StringSlice(matches)))

		for _, f := range matches[w.MaxCount:] {
			file := filepath.Join(dir, f)
			os.Remove(file)
		}
	}
}
