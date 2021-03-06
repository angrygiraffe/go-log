package log

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// BizFormatter is a text line formatter specifical for bizbase
type BizFormatter struct {
	AppName    string
	TimeFormat string

	init   sync.Once
	host   []byte
	app    []byte
	pid    []byte
	isterm bool
}

// Format implements Formatter
func (f *BizFormatter) Format(level Level, msg string, logger *Logger) []byte {
	// output format: DATE LEVEL HOST APP PID file:line message
	// 2001-10-10T12:00:00,000+0800 INFO web-1 app 1234 main/main.go:1234 message ...

	f.init.Do(func() {
		if f.AppName == "" {
			f.AppName = filepath.Base(os.Args[0])
		}
		f.app = []byte(f.AppName)

		if f.TimeFormat == "" {
			f.TimeFormat = "[2006-01-02 15:04:05.000]"
		}

		f.isterm = IsTerminal(logger.Out)

		host, _ := os.Hostname()
		f.host = []byte(host)

		f.pid = []byte(strconv.Itoa(os.Getpid()))
	})

	buf := fmtBuffer.Get().(*bytes.Buffer)
	buf.Reset()
	defer fmtBuffer.Put(buf)

	// timestamp
	timeStr := time.Now().Format(f.TimeFormat)
	buf.WriteString(timeStr)

	// level
	buf.WriteByte(' ')
	if f.isterm {
		buf.WriteString(level.ColorString())
	} else {
		buf.WriteString(level.String())
	}

	// host
	buf.WriteByte(' ')
	buf.Write(f.host)

	// name
	buf.WriteByte(' ')
	buf.Write(f.app)

	// pid
	buf.WriteByte(' ')
	buf.Write(f.pid)

	// file, line
	file, line := FilelineCaller(5)
	buf.WriteByte(' ')
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))

	// msg
	buf.WriteByte(' ')
	buf.WriteString(msg)

	// newline
	buf.WriteByte('\n')

	return buf.Bytes()
}
