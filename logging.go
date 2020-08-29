// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	ShortFile = 1
	LongFile  = 1 << 1
)

const (
	AutoColor = iota
	DisableColor
	ForceColor
)

var (
	ColorGreen   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	ColorWhite   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	ColorYellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	ColorRed     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	ColorBlue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	ColorMagenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	ColorCyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	ColorReset   = string([]byte{27, 91, 48, 109})

	ForeGreen   = "\033[97;32m"
	ForeWhite   = "\033[90;37m"
	ForeYellow  = "\033[90;33m"
	ForeRed     = "\033[97;31m"
	ForeBlue    = "\033[97;34m"
	ForeMagenta = "\033[97;35m"
	ForeCyan    = "\033[97;36m"

	BackGreen   = "\033[97;42m"
	BackWhite   = "\033[90;47m"
	BackYellow  = "\033[90;43m"
	BackRed     = "\033[97;41m"
	BackBlue    = "\033[97;44m"
	BackMagenta = "\033[97;45m"
	BackCyan    = "\033[97;46m"

	ResetColor = "\033[0m"
)

var gLogTag = map[int]string{
	DEBUG: "Debug",
	INFO:  "Info",
	WARN:  "Warn",
	ERROR: "Error",
	FATAL: "Fatal",
}

type LoggingOpt func(l *Logging)

type Logging struct {
	timeFormatter func(t time.Time) string
	formatter     func(writer io.Writer, level, depth int, tag string, info string)
	colorFlag     int
	fileFlag      int
	fatalNoTrace  bool

	level int

	writers map[int]io.Writer

	bufPool sync.Pool
}

var (
	ColorFlag     = AutoColor
	PrintFileFlag = ShortFile
	FatalNoTrace  = false
	Level         = INFO
	Writers       = map[int]io.Writer{
		DEBUG: os.Stdout,
		INFO:  os.Stdout,
		WARN:  os.Stdout,
		ERROR: os.Stderr,
		FATAL: os.Stderr,
	}
)

var defaultLogging = NewLogging()

func NewLogging(opts ...LoggingOpt) *Logging {
	ret := &Logging{
		timeFormatter: TimeFormat,
		colorFlag:     ColorFlag,
		fileFlag:      PrintFileFlag,
		fatalNoTrace:  FatalNoTrace,
		level:         Level,

		writers: Writers,

		bufPool: sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(nil)
		}},
	}
	ret.formatter = ret.format

	for _, v := range opts {
		v(ret)
	}
	return ret
}

func (l *Logging) format(writer io.Writer, level, depth int, tag string, log string) {
	var (
		file string
		line int
		ok   bool
	)
	_, file, line, ok = runtime.Caller(2 + depth)
	if !ok {
		file = "???"
		line = 0
	}

	var (
		lvColor    string
		resetColor string
	)
	if l.colorFlag == AutoColor {
		lvColor = selectLevelColor(level)
		resetColor = ResetColor
	}

	if l.fileFlag == 0 {
		file = ""
	} else if (l.fileFlag & ShortFile) != 0 {
		file = shortFile(file)
	}

	if tag != "" {
		fmt.Fprintf(writer, "%s [%s%s%s] [%s:%d] [%s] %s ",
			l.timeFormatter(time.Now()), lvColor, gLogTag[level], resetColor, file, line, tag, log)
	} else {
		fmt.Fprintf(writer, "%s [%s%s%s] [%s:%d] %s ",
			l.timeFormatter(time.Now()), lvColor, gLogTag[level], resetColor, file, line, log)
	}
}

func selectLevelColor(level int) string {
	if level == INFO {
		return ForeCyan
	} else if level == WARN {
		return ForeYellow
	} else if level > WARN {
		return ForeRed
	}
	return ""
}

func (l *Logging) Logf(level int, depth int, tag string, format string, args ...interface{}) {
	if l.level > level {
		return
	}

	length := len(format)
	if length > 0 {
		if format[length-1] != '\n' {
			format = format + "\n"
		}
	}
	logInfo := fmt.Sprintf(format, args...)
	w := l.selectWriter(level)
	l.formatter(w, level, depth, tag, logInfo)

	if level >= FATAL {
		l.processFatal(w)
	}

	//l.output(level, buf)
}

func (l *Logging) Log(level int, depth int, tag string, args ...interface{}) {
	if l.level > level {
		return
	}

	logInfo := fmt.Sprint(args...)
	w := l.selectWriter(level)
	l.formatter(w, level, depth, tag, logInfo)

	if level >= FATAL {
		l.processFatal(w)
	}
}

func (l *Logging) Logln(level int, depth int, tag string, args ...interface{}) {
	if l.level > level {
		return
	}

	logInfo := fmt.Sprintln(args...)
	w := l.selectWriter(level)
	l.formatter(w, level, depth, tag, logInfo)

	if level >= FATAL {
		l.processFatal(w)
	}
}

func (l *Logging) getBuffer() *bytes.Buffer {
	buf := l.bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (l *Logging) putBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	if buf.Len() > 256 {
		//let big buffers die a natural death.
		return
	}
	l.bufPool.Put(buf)
}

func (l *Logging) processFatal(writer io.Writer) {
	if !l.fatalNoTrace {
		trace := stacks(true)
		writer.Write(trace)
	}
	os.Exit(-1)
}

//func (l *Logging) output(level int) {
//	if level >= FATAL {
//		if !l.fatalNoTrace {
//			trace := stacks(true)
//			buf.WriteString("\n")
//			buf.Write(trace)
//		}
//		l.selectWriter(level).Write(buf.Bytes())
//		os.Exit(-1)
//	} else {
//		l.selectWriter(level).Write(buf.Bytes())
//	}
//	l.putBuffer(buf)
//}

func (l *Logging) selectWriter(level int) io.Writer {
	w := l.writers[level]
	if w == nil {
		return os.Stdout
	}
	return w
}

func shortFile(file string) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}

func stacks(all bool) []byte {
	n := 10000
	if all {
		n = 100000
	}
	var trace []byte
	for i := 0; i < 5; i++ {
		trace = make([]byte, n)
		nbytes := runtime.Stack(trace, all)
		if nbytes < len(trace) {
			return trace[:nbytes]
		}
	}
	return trace
}

func TimeFormat(t time.Time) string {
	var timeString = t.Format("2006-01-02 15:04:05")
	return timeString
}

func SetTimeFormatter(f func(t time.Time) string) func(*Logging) {
	return func(logging *Logging) {
		logging.timeFormatter = f
	}
}

func SetHeaderFormatter(f func(writer io.Writer, level, depth int, tag, log string)) func(*Logging) {
	return func(logging *Logging) {
		logging.formatter = f
	}
}

func SetColorFlag(flag int) func(*Logging) {
	return func(logging *Logging) {
		logging.colorFlag = flag
	}
}

func SetShowFileFlag(flag int) func(*Logging) {
	return func(logging *Logging) {
		logging.fileFlag = flag
	}
}

func SetFatalNoTrace(noTrace bool) func(*Logging) {
	return func(logging *Logging) {
		logging.fatalNoTrace = noTrace
	}
}

func SetLogSeverity(severity int) func(*Logging) {
	return func(logging *Logging) {
		logging.level = severity
	}
}

func SetOutput(w io.Writer) func(*Logging) {
	return func(logging *Logging) {
		for i := DEBUG; i <= FATAL; i++ {
			logging.writers[i] = w
		}
	}
}

func SetOutputBySeverity(severity int, w io.Writer) func(*Logging) {
	return func(logging *Logging) {
		logging.writers[severity] = w
	}
}

func Init(logging *Logging) {
	defaultLogging = logging
	ResetFactory(logging)
}

const autogeneratedFrameName = "<autogenerated>"

func FramesToCaller() int {
	for i := 1; i < 3; i++ {
		_, file, _, _ := runtime.Caller(i + 1)
		if file != autogeneratedFrameName {
			return i
		}
	}
	return 1
}
