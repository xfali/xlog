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

	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
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
	timeFormatter   func(t time.Time) string
	headerFormatter func(level, depth int) string
	colorFlag       int
	fileFlag        int
	fatalNoTrace    bool

	level int

	writers map[int]io.Writer
}

func NewLogging(opts ...LoggingOpt) *Logging {
	ret := &Logging{
		timeFormatter: TimeFormat,
		colorFlag:     AutoColor,
		fileFlag:      ShortFile,
		fatalNoTrace:  false,
		level:         INFO,

		writers: map[int]io.Writer{
			DEBUG: os.Stdout,
			INFO:  os.Stdout,
			WARN:  os.Stdout,
			ERROR: os.Stderr,
			FATAL: os.Stderr,
		},
	}
	ret.headerFormatter = ret.formatHeader

	for _, v := range opts {
		v(ret)
	}
	return ret
}

func (l *Logging) formatHeader(level, depth int) string {
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
		resetColor = ColorReset
	}

	if l.fileFlag == 0 {
		file = ""
	} else if (l.fileFlag & ShortFile) != 0 {
		file = shortFile(file)
	}

	return fmt.Sprintf("%s [%s%s%s] [%s:%d] ",
		l.timeFormatter(time.Now()), lvColor, gLogTag[level], resetColor, file, line)
}

func selectLevelColor(level int) string {
	if level == INFO {
		return blue
	} else if level == WARN {
		return yellow
	} else if level > WARN {
		return red
	}
	return ""
}

func (l *Logging) Logf(level int, depth int, format string, args ...interface{}) {
	if l.level > level {
		return
	}

	buf := bytes.NewBufferString(l.headerFormatter(level, depth))
	fmt.Fprintf(buf, format, args...)
	if buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}

	l.output(level, buf)
}

func (l *Logging) Log(level int, depth int, args ...interface{}) {
	if l.level > level {
		return
	}

	buf := bytes.NewBufferString(l.headerFormatter(level, depth))
	fmt.Fprint(buf, args...)

	l.output(level, buf)
}

func (l *Logging) Logln(level int, depth int, args ...interface{}) {
	if l.level > level {
		return
	}

	buf := bytes.NewBufferString(l.headerFormatter(level, depth))
	fmt.Fprintln(buf, args...)

	l.output(level, buf)
}

func (l *Logging) output(level int, buf *bytes.Buffer) {
	if level >= FATAL {
		if !l.fatalNoTrace {
			trace := stacks(true)
			buf.WriteString("\n")
			buf.Write(trace)
		}
		l.selectWriter(level).Write(buf.Bytes())
		os.Exit(-1)
	} else {
		l.selectWriter(level).Write(buf.Bytes())
	}
}

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

func SetHeaderFormatter(f func(level, depth int) string) func(*Logging) {
	return func(logging *Logging) {
		logging.headerFormatter = f
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
