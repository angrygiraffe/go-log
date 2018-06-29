package main

import (
	"os"

	"github.com/angrygiraffe/go-log"
	"github.com/angrygiraffe/go-log/formatters"
)

func main() {
	log.Info("hello", 123)
	log.Warn("hello", 123)

	log.Default.Formatter = new(formatters.BizFormatter)
	log.Infoln("hello", "world")
	log.Warnln("hello", "world")

	newLog := &log.Logger{
		Level:     log.INFO,
		Formatter: new(formatters.JSONFormatter),
		Out:       os.Stdout,
	}
	newLog.Infof("hello %v", 123)
	newLog.Warnf("hello %v", 123)
}
