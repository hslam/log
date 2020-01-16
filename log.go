// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package log implements multilevel logging.
package log

import (
	"io"
	"log"
	"os"
)

// Level defines the level for log.
// Higher levels log less info.
type Level int

const (
	//DebugLevel defines the level of debug in test environments.
	DebugLevel Level = 1
	//TraceLevel defines the level of trace in test environments.
	TraceLevel Level = 2
	//AllLevel defines the lowest level in production environments.
	AllLevel Level = 3
	//InfoLevel defines the level of info.
	InfoLevel Level = 4
	//NoticeLevel defines the level of notice.
	NoticeLevel Level = 5
	//WarnLevel defines the level of warn.
	WarnLevel Level = 6
	//ErrorLevel defines the level of error.
	ErrorLevel Level = 7
	//PanicLevel defines the level of panic.
	PanicLevel Level = 8
	//FatalLevel defines the level of fatal.
	FatalLevel Level = 9
	//OffLevel defines the level of no log.
	OffLevel Level = 10
)

var (
	out          io.Writer
	logPrefix    = "Log"
	logLevel     Level
	debugLogger  *log.Logger
	traceLogger  *log.Logger
	allLogger    *log.Logger
	infoLogger   *log.Logger
	noticeLogger *log.Logger
	warnLogger   *log.Logger
	errorLogger  *log.Logger
	panicLogger  *log.Logger
	fatalLogger  *log.Logger
	redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	black        = string([]byte{27, 91, 57, 48, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	magenta      = string([]byte{27, 91, 51, 53, 109})
	cyan         = string([]byte{27, 91, 51, 54, 109})
	white        = string([]byte{27, 91, 51, 55, 109})
	reset        = string([]byte{27, 91, 48, 109})
)

func init() {
	SetOut(os.Stdout)
	SetLevel(InfoLevel)
	initLog()
}

func initLog() {
	debugLogger = log.New(out, blue+"["+logPrefix+"][D]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	traceLogger = log.New(out, cyan+"["+logPrefix+"][T]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	allLogger = log.New(out, white+"["+logPrefix+"][A]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	infoLogger = log.New(out, black+"["+logPrefix+"][I]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	noticeLogger = log.New(out, green+"["+logPrefix+"][N]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	warnLogger = log.New(out, yellow+"["+logPrefix+"][W]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	errorLogger = log.New(out, redBg+"["+logPrefix+"][E]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	panicLogger = log.New(out, red+"["+logPrefix+"][P]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	fatalLogger = log.New(out, magenta+"["+logPrefix+"][F]"+reset, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
}

//SetPrefix sets log's prefix
func SetPrefix(prefix string) {
	logPrefix = prefix
	initLog()
}

//SetLevel sets log's level
func SetLevel(level Level) {
	logLevel = level
}

//SetOut sets log's writer. The out variable sets the
// destination to which log data will be written.
func SetOut(w io.Writer) {
	out = w
}

//GetLevel returns log's level
func GetLevel() Level {
	return logLevel
}

// Debug is equivalent to log.Print() for debug.
func Debug(v ...interface{}) {
	if logLevel <= DebugLevel {
		debugLogger.Print(v...)
	}
}

// Debugf is equivalent to log.Printf() for debug.
func Debugf(format string, v ...interface{}) {
	if logLevel <= DebugLevel {
		debugLogger.Printf(format, v...)
	}
}

// Debugln is equivalent to log.Println() for debug.
func Debugln(v ...interface{}) {
	if logLevel <= DebugLevel {
		debugLogger.Println(v...)
	}
}

// Trace is equivalent to log.Print() for trace.
func Trace(v ...interface{}) {
	if logLevel <= TraceLevel {
		traceLogger.Print(v...)
	}
}

// Tracef is equivalent to log.Printf() for trace.
func Tracef(format string, v ...interface{}) {
	if logLevel <= TraceLevel {
		traceLogger.Printf(format, v...)
	}
}

// Traceln is equivalent to log.Println() for trace.
func Traceln(v ...interface{}) {
	if logLevel <= TraceLevel {
		traceLogger.Println(v...)
	}
}

// All is equivalent to log.Print() for all log.
func All(v ...interface{}) {
	if logLevel <= AllLevel {
		allLogger.Print(v...)
	}
}

// Allf is equivalent to log.Printf() for all log.
func Allf(format string, v ...interface{}) {
	if logLevel <= AllLevel {
		allLogger.Printf(format, v...)
	}
}

// Allln is equivalent to log.Println() for all log.
func Allln(v ...interface{}) {
	if logLevel <= AllLevel {
		allLogger.Println(v...)
	}
}

// Info is equivalent to log.Print() for info.
func Info(v ...interface{}) {
	if logLevel <= InfoLevel {
		infoLogger.Print(v...)
	}
}

// Infof is equivalent to log.Printf() for info.
func Infof(format string, v ...interface{}) {
	if logLevel <= InfoLevel {
		infoLogger.Printf(format, v...)
	}
}

// Infoln is equivalent to log.Println() for info.
func Infoln(v ...interface{}) {
	if logLevel <= InfoLevel {
		infoLogger.Println(v...)
	}
}

// Notice is equivalent to log.Print() for notice.
func Notice(v ...interface{}) {
	if logLevel <= NoticeLevel {
		noticeLogger.Print(v...)
	}
}

// Noticef is equivalent to log.Printf() for notice.
func Noticef(format string, v ...interface{}) {
	if logLevel <= NoticeLevel {
		noticeLogger.Printf(format, v...)
	}
}

// Noticeln is equivalent to log.Println() for notice.
func Noticeln(v ...interface{}) {
	if logLevel <= NoticeLevel {
		noticeLogger.Println(v...)
	}
}

// Warn is equivalent to log.Print() for warn.
func Warn(v ...interface{}) {
	if logLevel <= WarnLevel {
		warnLogger.Print(v...)
	}
}

// Warnf is equivalent to log.Printf() for warn.
func Warnf(format string, v ...interface{}) {
	if logLevel <= InfoLevel {
		warnLogger.Printf(format, v...)
	}
}

// Warnln is equivalent to log.Println() for warn.
func Warnln(v ...interface{}) {
	if logLevel <= WarnLevel {
		warnLogger.Println(v...)
	}
}

// Error is equivalent to log.Print() for error.
func Error(v ...interface{}) {
	if logLevel <= ErrorLevel {
		errorLogger.Print(v...)
	}
}

// Errorf is equivalent to log.Printf() for error.
func Errorf(format string, v ...interface{}) {
	if logLevel <= ErrorLevel {
		errorLogger.Printf(format, v...)
	}
}

// Errorln is equivalent to log.Println() for error.
func Errorln(v ...interface{}) {
	if logLevel <= ErrorLevel {
		errorLogger.Println(v...)
	}
}

// Panic is equivalent to log.Print() for panic.
func Panic(v ...interface{}) {
	if logLevel <= PanicLevel {
		panicLogger.Print(v...)
	}
}

// Panicf is equivalent to log.Printf() for panic.
func Panicf(format string, v ...interface{}) {
	if logLevel <= PanicLevel {
		panicLogger.Printf(format, v...)
	}
}

// Panicln is equivalent to log.Println() for panic.
func Panicln(v ...interface{}) {
	if logLevel <= PanicLevel {
		panicLogger.Println(v...)
	}
}

// Fatal is equivalent to log.Print() for fatal.
func Fatal(v ...interface{}) {
	if logLevel <= FatalLevel {
		fatalLogger.Print(v...)
	}
}

// Fatalf is equivalent to log.Printf() for fatal.
func Fatalf(format string, v ...interface{}) {
	if logLevel <= FatalLevel {
		fatalLogger.Printf(format, v...)
	}
}

// Fatalln is equivalent to log.Println() for fatal.
func Fatalln(v ...interface{}) {
	if logLevel <= FatalLevel {
		fatalLogger.Println(v...)
	}
}
