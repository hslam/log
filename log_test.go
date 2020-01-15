// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package log implements logging of multi level.
package log

import (
	"testing"
)

func TestSetPrefix(t *testing.T) {
	var prefix = "log"
	SetPrefix(prefix)
	if prefix != logPrefix {
		t.Errorf("error %s != %s", prefix, logPrefix)
	}
}

func TestSetLevel(t *testing.T) {
	var level = DebugLevel
	SetLevel(level)
	if level != GetLevel() {
		t.Errorf("error %d != %d", level, GetLevel())
	}
}

func TestDebug(t *testing.T) {
	Debug(1024, "HelloWorld", true)
	Debugf("%d %s %t", 1024, "HelloWorld", true)
	Debugln(1024, "HelloWorld", true)
}

func TestTrace(t *testing.T) {
	Trace(1024, "HelloWorld", true)
	Tracef("%d %s %t", 1024, "HelloWorld", true)
	Traceln(1024, "HelloWorld", true)
}

func TestAll(t *testing.T) {
	All(1024, "HelloWorld", true)
	Allf("%d %s %t", 1024, "HelloWorld", true)
	Allln(1024, "HelloWorld", true)
}

func TestInfo(t *testing.T) {
	Info(1024, "HelloWorld", true)
	Infof("%d %s %t", 1024, "HelloWorld", true)
	Infoln(1024, "HelloWorld", true)
}

func TestWarn(t *testing.T) {
	Warn(1024, "HelloWorld", true)
	Warnf("%d %s %t", 1024, "HelloWorld", true)
	Warnln(1024, "HelloWorld", true)
}

func TestError(t *testing.T) {
	Error(1024, "HelloWorld", true)
	Errorf("%d %s %t", 1024, "HelloWorld", true)
	Errorln(1024, "HelloWorld", true)
}

func TestPanic(t *testing.T) {
	Panic(1024, "HelloWorld", true)
	Panicf("%d %s %t", 1024, "HelloWorld", true)
	Panicln(1024, "HelloWorld", true)
}

func TestFatal(t *testing.T) {
	Fatal(1024, "HelloWorld", true)
	Fatalf("%d %s %t", 1024, "HelloWorld", true)
	Fatalln(1024, "HelloWorld", true)
}

func TestOff(t *testing.T) {
	Off(1024, "HelloWorld", true)
	Offf("%d %s %t", 1024, "HelloWorld", true)
	Offln(1024, "HelloWorld", true)
}
