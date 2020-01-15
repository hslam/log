// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package log implements multilevel logging.
package log

import (
	"log"
	"os"
)

// Level defines the level for log.
// Higher levels log less info.
type Level int

const (

	//DebugLevel defines the level of debug.
	DebugLevel Level = 1
	//TraceLevel defines the level of trace.
	TraceLevel Level = 2
	//AllLevel defines the level of all.
	AllLevel Level = 3
	//InfoLevel defines the level of info.
	InfoLevel Level = 4
	//WarnLevel defines the level of warn.
	WarnLevel Level = 5
	//ErrorLevel defines the level of error.
	ErrorLevel Level = 6
	//PanicLevel defines the level of panic.
	PanicLevel Level = 7
	//FatalLevel defines the level of fatal.
	FatalLevel Level = 8
	//OffLevel defines the level of off.
	OffLevel Level = 9
	//NoLevel defines the level of no log.
	NoLevel Level = 99
)

var (
	logPrefix   = "Log"
	logLevel    Level
	debugLogger *log.Logger
	traceLogger *log.Logger
	allLogger   *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	panicLogger *log.Logger
	fatalLogger *log.Logger
	offLogger   *log.Logger
)

func init() {
	SetLevel(AllLevel)
	initLog()
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

func initLog() {
	debugLogger = log.New(os.Stdout, "["+logPrefix+"][D] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	traceLogger = log.New(os.Stdout, "["+logPrefix+"][T] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	allLogger = log.New(os.Stdout, "["+logPrefix+"][A] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	infoLogger = log.New(os.Stdout, "["+logPrefix+"][I] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	warnLogger = log.New(os.Stdout, "["+logPrefix+"][W] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	errorLogger = log.New(os.Stdout, "["+logPrefix+"][E] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	panicLogger = log.New(os.Stdout, "["+logPrefix+"][P] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	fatalLogger = log.New(os.Stdout, "["+logPrefix+"][F] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	offLogger = log.New(os.Stdout, "["+logPrefix+"][O] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)

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

// Off is equivalent to log.Print() for off.
func Off(v ...interface{}) {
	if logLevel <= OffLevel {
		offLogger.Print(v...)
	}
}

// Offf is equivalent to log.Printf() for off.
func Offf(format string, v ...interface{}) {
	if logLevel <= OffLevel {
		offLogger.Printf(format, v...)
	}
}

// Offln is equivalent to log.Println() for off.
func Offln(v ...interface{}) {
	if logLevel <= OffLevel {
		offLogger.Println(v...)
	}
}
