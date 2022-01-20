// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package log implements multilevel logging.
package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level defines the level for log.
// Higher levels log less info.
type Level uint8

const (
	//DebugLevel defines the level of debug in test environments.
	DebugLevel Level = 0
	//TraceLevel defines the level of trace in test environments.
	TraceLevel Level = 1
	//AllLevel defines the lowest level in production environments.
	AllLevel Level = 2
	//InfoLevel defines the level of info.
	InfoLevel Level = 3
	//NoticeLevel defines the level of notice.
	NoticeLevel Level = 4
	//WarnLevel defines the level of warn.
	WarnLevel Level = 5
	//ErrorLevel defines the level of error.
	ErrorLevel Level = 6
	//PanicLevel defines the level of panic.
	PanicLevel Level = 7
	//FatalLevel defines the level of fatal.
	FatalLevel Level = 8
	//OffLevel defines the level of no log.
	OffLevel Level = 9
)

const (
	ignorePackagePrefix = "github.com/hslam/log."
	logPrefix           = "Log"
)

var (
	logger    = New()
	bufPool   sync.Pool
	pcPool    sync.Pool
	bigpcPool sync.Pool
)

// colors
var (
	redBg     = []byte{27, 91, 57, 55, 59, 52, 49, 109}
	magentaBg = []byte{27, 91, 57, 55, 59, 52, 53, 109}
	black     = []byte{27, 91, 57, 48, 109}
	red       = []byte{27, 91, 51, 49, 109}
	green     = []byte{27, 91, 51, 50, 109}
	yellow    = []byte{27, 91, 51, 51, 109}
	blue      = []byte{27, 91, 51, 52, 109}
	magenta   = []byte{27, 91, 51, 53, 109}
	cyan      = []byte{27, 91, 51, 54, 109}
	white     = []byte{27, 91, 51, 55, 109}
	reset     = []byte{27, 91, 48, 109}
)

// newBuffer returns a new buffer writer.
func newBuffer() *bytes.Buffer {
	if v := bufPool.Get(); v != nil {
		buf := v.(*bytes.Buffer)
		return buf
	}
	return bytes.NewBuffer(nil)
}

// freeBuffer frees the buffer writer.
func freeBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufPool.Put(buf)
}

const pcs = 16

func newPC() []uintptr {
	if v := pcPool.Get(); v != nil {
		pc := v.([]uintptr)
		return pc
	}
	return make([]uintptr, pcs)
}

func freePC(pc []uintptr) {
	if cap(pc) == pcs {
		pcPool.Put(pc[:pcs])
	}
}

const bigpcs = 4096

func newBigPC() []uintptr {
	if v := bigpcPool.Get(); v != nil {
		pc := v.([]uintptr)
		return pc
	}
	return make([]uintptr, bigpcs)
}

func freeBigPC(pc []uintptr) {
	if cap(pc) == bigpcs {
		bigpcPool.Put(pc[:bigpcs])
	}
}

// relevantCaller searches the call stack for the first function outside of log.
// The purpose of this function is to provide more helpful error messages.
func relevantCaller() runtime.Frame {
	pc := newPC()
	n := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:n])
	var frame runtime.Frame
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, ignorePackagePrefix) {
			freePC(pc)
			return frame
		}
		if !more {
			break
		}
	}
	freePC(pc)
	return frame
}

const (
	newline     = "\\n"
	frameFormat = "%s\\n\\t%s:%d"
)

func callStack() (s string) {
	var stack []string
	pc := newBigPC()
	n := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, ignorePackagePrefix) {
			stack = append(stack, fmt.Sprintf(frameFormat, frame.Function, frame.File, frame.Line))
		}
		if !more {
			break
		}
	}
	if len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}
	freeBigPC(pc)
	return strings.Join(stack, newline)
}

// log defines the base log interface.
type log interface {
	// Output writes the log info to the io.Writer.
	Output(w io.Writer, b []byte)
}

// body implements the log interface.
type body struct {
}

// newBody returns a new log body.
func newBody() log {
	return body{}
}

var (
	front = []byte("[")
	back  = []byte("]")
)

// Output writes the log info to the io.Writer.
func (l body) Output(w io.Writer, b []byte) {
	w.Write(front)
	w.Write(b)
	w.Write(back)
}

// stackField implements the log interface.
type stackField struct {
	l log
}

// withStackField returns a new log with the stack field.
func withStackField(l log) log {
	return &stackField{l}
}

const stackFormat = "[stack=\"%s\"]"

// Output writes the log info to the io.Writer.
func (l *stackField) Output(w io.Writer, b []byte) {
	l.l.Output(w, b)
	fmt.Fprintf(w, stackFormat, callStack())
}

// lineField implements the log interface.
type lineField struct {
	l log
}

// withLineField returns a new log with the line field.
func withLineField(l log) log {
	return &lineField{l}
}

const callerFormat = "[%s:%d]"

// Output writes the log info to the io.Writer.
func (l *lineField) Output(w io.Writer, b []byte) {
	caller := relevantCaller()
	fmt.Fprintf(w, callerFormat, path.Base(caller.File), caller.Line)
	l.l.Output(w, b)
}

// levelField implements the log interface.
type levelField struct {
	l     log
	level []byte
}

// withLevelField returns a new log with the level field.
func withLevelField(l log, level string) log {
	return &levelField{l, []byte("[" + level + "]")}
}

// Output writes the log info to the io.Writer.
func (l *levelField) Output(w io.Writer, b []byte) {
	w.Write(l.level)
	l.l.Output(w, b)
}

// timeField implements the log interface.
type timeField struct {
	l log
}

// withTimeField returns a new log with the time field.
func withTimeField(l log) log {
	return &timeField{l}
}

const (
	timeFormat    = "[2006/01/02 15:04:05.000 -07:00]"
	timeFormatLen = len(timeFormat)
)

// Output writes the log info to the io.Writer.
func (l *timeField) Output(w io.Writer, b []byte) {
	var buf [timeFormatLen]byte
	var tb = buf[:0]
	time.Now().AppendFormat(tb, timeFormat)
	w.Write(tb[:timeFormatLen])
	l.l.Output(w, b)
}

// prefixField implements the log interface.
type prefixField struct {
	l    log
	name []byte
}

// withPrefixField returns a new log with the prefix field.
func withPrefixField(l log, prefix string) log {
	return &prefixField{l, []byte("[" + prefix + "]")}
}

// Output writes the log info to the io.Writer.
func (l *prefixField) Output(w io.Writer, b []byte) {
	w.Write(l.name)
	l.l.Output(w, b)
}

// highlightField implements the log interface.
type highlightField struct {
	l     log
	color []byte
}

// withHighlightField returns a new log with the highlight field.
func withHighlightField(l log, color []byte) log {
	return &highlightField{l, color}
}

// Output writes the log info to the io.Writer.
func (l *highlightField) Output(w io.Writer, b []byte) {
	w.Write(l.color)
	l.l.Output(w, b)
	w.Write(reset)
}

var levels = [9]string{"D", "T", "A", "I", "N", "W", "E", "P", "F"}
var colors = [9][]byte{blue, cyan, white, black, green, yellow, red, magentaBg, redBg}

func newLog(prefix string, level Level, highlight, line bool) log {
	l := newBody()
	if line {
		l = withLineField(l)
	}
	if level >= ErrorLevel || level == TraceLevel {
		l = withStackField(l)
	}
	l = withLevelField(l, levels[level])
	l = withTimeField(l)
	if len(prefix) > 0 {
		l = withPrefixField(l, prefix)
	}
	color := colors[level]
	if len(color) > 0 && highlight {
		l = withHighlightField(l, color)
	}
	return l
}

// Logger defines the logger.
type Logger struct {
	mu        sync.Mutex
	out       io.Writer
	prefix    string
	level     Level
	highlight bool
	line      bool
	logs      [9]log
}

// New creates a new Logger.
func New() *Logger {
	l := &Logger{
		out:    os.Stdout,
		prefix: logPrefix,
		level:  InfoLevel,
		line:   true,
	}
	l.init()
	return l
}

//SetPrefix sets log's prefix
func SetPrefix(prefix string) {
	logger.SetPrefix(prefix)
}

//SetPrefix sets log's prefix
func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
	l.init()
}

//GetPrefix returns log's prefix
func GetPrefix() (prefix string) {
	return logger.GetPrefix()
}

//GetPrefix returns log's prefix
func (l *Logger) GetPrefix() (prefix string) {
	return l.prefix
}

//SetLevel sets log's level
func SetLevel(level Level) {
	logger.SetLevel(level)
}

//SetLevel sets log's level
func (l *Logger) SetLevel(level Level) {
	l.level = level
	l.init()
}

//SetHighlight sets whether to enable the highlight field.
func SetHighlight(highlight bool) {
	logger.SetHighlight(highlight)
}

//SetHighlight sets whether to enable the highlight field.
func (l *Logger) SetHighlight(highlight bool) {
	l.highlight = highlight
	l.init()
}

//SetLine sets whether to enable the line field .
func SetLine(line bool) {
	logger.SetLine(line)
}

//SetLine sets whether to enable the line field .
func (l *Logger) SetLine(line bool) {
	l.line = line
	l.init()
}

//SetOut sets log's writer. The out variable sets the
// destination to which log data will be written.
func SetOut(w io.Writer) {
	logger.SetOut(w)
}

//SetOut sets log's writer. The out variable sets the
// destination to which log data will be written.
func (l *Logger) SetOut(w io.Writer) {
	l.out = w
	l.init()
}

//GetLevel returns log's level
func GetLevel() Level {
	return logger.GetLevel()
}

//GetLevel returns log's level
func (l *Logger) GetLevel() Level {
	return l.level
}

func (l *Logger) init() {
	for i := 0; i < 9; i++ {
		l.logs[i] = newLog(l.prefix, Level(i), l.highlight, l.line)
	}
}

func (l *Logger) logout(level Level, b []byte) {
	buf := newBuffer()
	l.logs[level].Output(buf, bytes.TrimSpace(b))
	fmt.Fprintln(buf)
	l.mu.Lock()
	l.out.Write(buf.Bytes())
	l.mu.Unlock()
	freeBuffer(buf)
}

func (l *Logger) print(level Level, v ...interface{}) {
	body := newBuffer()
	fmt.Fprint(body, v...)
	l.logout(level, body.Bytes())
	freeBuffer(body)
}

func (l *Logger) printf(level Level, format string, v ...interface{}) {
	body := newBuffer()
	fmt.Fprintf(body, format, v...)
	l.logout(level, body.Bytes())
	freeBuffer(body)
}

func (l *Logger) println(level Level, v ...interface{}) {
	body := newBuffer()
	fmt.Fprintln(body, v...)
	l.logout(level, body.Bytes())
	freeBuffer(body)
}

// Debug is equivalent to log.Print() for debug.
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DebugLevel {
		l.print(DebugLevel, v...)
	}
}

// Debugf is equivalent to log.Printf() for debug.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= DebugLevel {
		l.printf(DebugLevel, format, v...)
	}
}

// Debugln is equivalent to log.Println() for debug.
func (l *Logger) Debugln(v ...interface{}) {
	if l.level <= DebugLevel {
		l.println(DebugLevel, v...)
	}
}

// Trace is equivalent to log.Print() for trace.
func (l *Logger) Trace(v ...interface{}) {
	if l.level <= TraceLevel {
		l.print(TraceLevel, v...)
	}
}

// Tracef is equivalent to log.Printf() for trace.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.level <= TraceLevel {
		l.printf(TraceLevel, format, v...)
	}
}

// Traceln is equivalent to log.Println() for trace.
func (l *Logger) Traceln(v ...interface{}) {
	if l.level <= TraceLevel {
		l.println(TraceLevel, v...)
	}
}

// All is equivalent to log.Print() for all log.
func (l *Logger) All(v ...interface{}) {
	if l.level <= AllLevel {
		l.print(AllLevel, v...)
	}
}

// Allf is equivalent to log.Printf() for all log.
func (l *Logger) Allf(format string, v ...interface{}) {
	if l.level <= AllLevel {
		l.printf(AllLevel, format, v...)
	}
}

// Allln is equivalent to log.Println() for all log.
func (l *Logger) Allln(v ...interface{}) {
	if l.level <= AllLevel {
		l.println(AllLevel, v...)
	}
}

// Info is equivalent to log.Print() for info.
func (l *Logger) Info(v ...interface{}) {
	if l.level <= InfoLevel {
		l.print(InfoLevel, v...)
	}
}

// Infof is equivalent to log.Printf() for info.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= InfoLevel {
		l.printf(InfoLevel, format, v...)
	}
}

// Infoln is equivalent to log.Println() for info.
func (l *Logger) Infoln(v ...interface{}) {
	if l.level <= InfoLevel {
		l.println(InfoLevel, v...)
	}
}

// Notice is equivalent to log.Print() for notice.
func (l *Logger) Notice(v ...interface{}) {
	if l.level <= NoticeLevel {
		l.print(NoticeLevel, v...)
	}
}

// Noticef is equivalent to log.Printf() for notice.
func (l *Logger) Noticef(format string, v ...interface{}) {
	if l.level <= NoticeLevel {
		l.printf(NoticeLevel, format, v...)
	}
}

// Noticeln is equivalent to log.Println() for notice.
func (l *Logger) Noticeln(v ...interface{}) {
	if l.level <= NoticeLevel {
		l.println(NoticeLevel, v...)
	}
}

// Warn is equivalent to log.Print() for warn.
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WarnLevel {
		l.print(WarnLevel, v...)
	}
}

// Warnf is equivalent to log.Printf() for warn.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level <= InfoLevel {
		l.printf(WarnLevel, format, v...)
	}
}

// Warnln is equivalent to log.Println() for warn.
func (l *Logger) Warnln(v ...interface{}) {
	if l.level <= WarnLevel {
		l.println(WarnLevel, v...)
	}
}

// Error is equivalent to log.Print() for error.
func (l *Logger) Error(v ...interface{}) {
	if l.level <= ErrorLevel {
		l.print(ErrorLevel, v...)
	}
}

// Errorf is equivalent to log.Printf() for error.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= ErrorLevel {
		l.printf(ErrorLevel, format, v...)
	}
}

// Errorln is equivalent to log.Println() for error.
func (l *Logger) Errorln(v ...interface{}) {
	if l.level <= ErrorLevel {
		l.println(ErrorLevel, v...)
	}
}

// Panic is equivalent to log.Print() for panic.
func (l *Logger) Panic(v ...interface{}) {
	if l.level <= PanicLevel {
		l.print(PanicLevel, v...)
		panic(fmt.Sprint(v...))
	}
}

// Panicf is equivalent to log.Printf() for panic.
func (l *Logger) Panicf(format string, v ...interface{}) {
	if l.level <= PanicLevel {
		l.printf(PanicLevel, format, v...)
		panic(fmt.Sprintf(format, v...))
	}
}

// Panicln is equivalent to log.Println() for panic.
func (l *Logger) Panicln(v ...interface{}) {
	if l.level <= PanicLevel {
		l.println(PanicLevel, v...)
		panic(fmt.Sprintln(v...))
	}
}

// Fatal is equivalent to log.Print() for fatal.
func (l *Logger) Fatal(v ...interface{}) {
	if l.level <= FatalLevel {
		l.print(FatalLevel, v...)
		os.Exit(1)
	}
}

// Fatalf is equivalent to log.Printf() for fatal.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level <= FatalLevel {
		l.printf(FatalLevel, format, v...)
		os.Exit(1)
	}
}

// Fatalln is equivalent to log.Println() for fatal.
func (l *Logger) Fatalln(v ...interface{}) {
	if l.level <= FatalLevel {
		l.println(FatalLevel, v...)
		os.Exit(1)
	}
}

// Debug is equivalent to log.Print() for debug.
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf is equivalent to log.Printf() for debug.
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Debugln is equivalent to log.Println() for debug.
func Debugln(v ...interface{}) {
	logger.Debugln(v...)
}

// Trace is equivalent to log.Print() for trace.
func Trace(v ...interface{}) {
	logger.Trace(v...)
}

// Tracef is equivalent to log.Printf() for trace.
func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

// Traceln is equivalent to log.Println() for trace.
func Traceln(v ...interface{}) {
	logger.Traceln(v...)
}

// All is equivalent to log.Print() for all log.
func All(v ...interface{}) {
	logger.All(v...)
}

// Allf is equivalent to log.Printf() for all log.
func Allf(format string, v ...interface{}) {
	logger.Allf(format, v...)
}

// Allln is equivalent to log.Println() for all log.
func Allln(v ...interface{}) {
	logger.Allln(v...)
}

// Info is equivalent to log.Print() for info.
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof is equivalent to log.Printf() for info.
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Infoln is equivalent to log.Println() for info.
func Infoln(v ...interface{}) {
	logger.Infoln(v...)
}

// Notice is equivalent to log.Print() for notice.
func Notice(v ...interface{}) {
	logger.Notice(v...)
}

// Noticef is equivalent to log.Printf() for notice.
func Noticef(format string, v ...interface{}) {
	logger.Noticef(format, v...)
}

// Noticeln is equivalent to log.Println() for notice.
func Noticeln(v ...interface{}) {
	logger.Noticeln(v...)
}

// Warn is equivalent to log.Print() for warn.
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Warnf is equivalent to log.Printf() for warn.
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Warnln is equivalent to log.Println() for warn.
func Warnln(v ...interface{}) {
	logger.Warnln(v...)
}

// Error is equivalent to log.Print() for error.
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf is equivalent to log.Printf() for error.
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Errorln is equivalent to log.Println() for error.
func Errorln(v ...interface{}) {
	logger.Errorln(v...)
}

// Panic is equivalent to log.Print() for panic.
func Panic(v ...interface{}) {
	logger.Panic(v...)
}

// Panicf is equivalent to log.Printf() for panic.
func Panicf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

// Panicln is equivalent to log.Println() for panic.
func Panicln(v ...interface{}) {
	logger.Panicln(v...)
}

// Fatal is equivalent to log.Print() for fatal.
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf is equivalent to log.Printf() for fatal.
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

// Fatalln is equivalent to log.Println() for fatal.
func Fatalln(v ...interface{}) {
	logger.Fatalln(v...)
}
