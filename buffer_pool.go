package log

import (
	"bytes"
	"sync"
)

var (
	fmtBuffer = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)
