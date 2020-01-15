package main

import (
	"github.com/hslam/log"
)

func main() {
	log.SetPrefix("LogPrefix")
	log.SetLevel(log.DebugLevel)
	log.Debugf("%d %s %t", 1024, "HelloWorld", true)
	log.Tracef("%d %s %t", 1024, "HelloWorld", true)
	log.Allf("%d %s %t", 1024, "HelloWorld", true)
	log.Infof("%d %s %t", 1024, "HelloWorld", true)
	log.Warnf("%d %s %t", 1024, "HelloWorld", true)
	log.Errorf("%d %s %t", 1024, "HelloWorld", true)
	log.Panicf("%d %s %t", 1024, "HelloWorld", true)
	log.Fatalf("%d %s %t", 1024, "HelloWorld", true)
	log.Offf("%d %s %t", 1024, "HelloWorld", true)
}
