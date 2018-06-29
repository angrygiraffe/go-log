package main

import (
	"github.com/subchen/go-log"
	"github.com/subchen/go-log/formatters"
	"github.com/subchen/go-log/writers"
	"time"
)

func main() {
	log.Info("hello", 123)
	log.Warn("hello", 123)

	log.Default.Formatter = new(formatters.BizFormatter)
	log.Infoln("hello", "world")
	log.Warnln("hello", "world")

	newLog := &log.Logger{
		Level:     log.INFO,
		Formatter: new(formatters.BizFormatter),
		Out: &writers.DailyFileWriter{
			Name:     "/tmp/test.log",
			MaxCount: 5,
			MaxSize:  1000,
		},
	}

	for i := 0; i < 1000; i++ {
		newLog.Infof("hello %v ................................", i)
		newLog.Warnf("bye   %v ................................", i)
		time.Sleep(time.Millisecond * 100)
	}
}
