package main

import (
	"github.com/hslam/log"
)

func main() {
	logger := log.New()
	logger.SetPrefix("LogPrefix")
	logger.SetLevel(log.DebugLevel)
	logger.Debugf("%d %s %t", 1024, "HelloWorld", true)
	logger.Tracef("%d %s %t", 1024, "HelloWorld", true)
	logger.Allf("%d %s %t", 1024, "HelloWorld", true)
	logger.Infof("%d %s %t", 1024, "HelloWorld", true)
	logger.Noticef("%d %s %t", 1024, "HelloWorld", true)
	logger.Warnf("%d %s %t", 1024, "HelloWorld", true)
	logger.Errorf("%d %s %t", 1024, "HelloWorld", true)
	logger.Panicf("%d %s %t", 1024, "HelloWorld", true)
	logger.Fatalf("%d %s %t", 1024, "HelloWorld", true)
}
