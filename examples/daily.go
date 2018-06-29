package main

import (
	"github.com/angrygiraffe/go-log"
	"time"
)

func main() {
	log.Info("hello", 123)
	log.Warn("hello", 123)

	log.Default.Formatter = new(log.BizFormatter)
	log.Infoln("hello", "world")
	log.Warnln("hello", "world")

	newLog := &log.Logger{
		Level:     log.INFO,
		Formatter: new(log.BizFormatter),
		Out: &log.DailyFileWriter{
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
