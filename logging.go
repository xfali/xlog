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
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Level = int32

const (
	DEBUG Level = 0
	INFO  Level = 1
	WARN  Level = 2
	ERROR Level = 3
	PANIC Level = 4
	FATAL Level = 5
)

const (
	//不显示调用信息
	CallerNone = 0
	//只显示文件名
	CallerShortFile = 1
	//显示文件名及路径
	CallerLongFile = 1 << 1
	//file mask
	CallerFileMask = 3
	//只显示函数名
	CallerShortFunc = 1 << 2
	//显示完整函数名
	CallerLongFunc = 1 << 3
	//func mask
	CallerFuncMask = 3 << 2
)

const (
	//自动填充颜色
	AutoColor = iota
	//禁用颜色
	DisableColor
	//强制使用颜色
	ForceColor
)

const (
	// 时间戳的Key
	KeyTimestamp = "LogTime"
	// 日志级别key
	KeySeverityLevel = "LogLevel"
	// 调用者Key
	KeyCaller = "LogCaller"
	// 日志内容Key
	KeyContent = "LogContent"
	// 日志名称Key
	KeyName = "LogName"
)

var (
	//前景色
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

	//背景色
	BackGreen   = "\033[97;42m"
	BackWhite   = "\033[90;47m"
	BackYellow  = "\033[90;43m"
	BackRed     = "\033[97;41m"
	BackBlue    = "\033[97;44m"
	BackMagenta = "\033[97;45m"
	BackCyan    = "\033[97;46m"

	ResetColor = "\033[0m"
)

// 级别及名称映射
var gLogTag = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	PANIC: "PANIC",
	FATAL: "FATAL",
}

// 默认值
var (
	DefaultColorFlag     = DisableColor
	DefaultPrintFileFlag = CallerShortFile
	DefaultFatalNoTrace  = false
	DefaultLevel         = INFO
	DefaultWriters       = map[Level]io.Writer{
		DEBUG: os.Stdout,
		INFO:  os.Stdout,
		WARN:  os.Stdout,
		ERROR: os.Stderr,
		PANIC: os.Stderr,
		FATAL: os.Stderr,
	}
)

type LoggingOpt func(l *logging)

// Logging是xlog的日志基础工具，向下对接日志输出Writer，向上提供日志操作接口
type Logging interface {
	// 输出format日志（线程安全）
	// Param： level日志级别， depth调用深度， keyValues附加的日志内容(多用于添加固定的日志信息)， format格式化的格式， args参数
	Logf(level Level, depth int, keyValues KeyValues, format string, args ...interface{})

	// 解析并输出参数（线程安全）
	// Param： level日志级别， depth调用深度， keyValues附加的日志内容(多用于添加固定的日志信息)， args参数
	Log(level Level, depth int, keyValues KeyValues, args ...interface{})

	// 解析并输出参数，末尾增加换行（线程安全）
	// Param： level日志级别， depth调用深度， keyValues附加的日志内容(多用于添加固定的日志信息)， args参数
	Logln(level Level, depth int, keyValues KeyValues, args ...interface{})

	// 设置日志格式化工具（线程安全）
	SetFormatter(f Formatter)

	// 设置日志严重级别，低于该级别的将不被输出（线程安全）
	SetSeverityLevel(severityLevel Level)

	// 判断参数级别是否会输出（线程安全）
	IsEnable(severityLevel Level) bool

	// 设置输出的Writer，注意该方法会将所有级别都配置为参数writer（线程安全）
	SetOutput(w io.Writer)

	// 设置对应日志级别的Writer（线程安全）
	SetOutputBySeverity(severityLevel Level, w io.Writer)

	// 获得对应日志级别的Writer（线程安全）
	GetOutputBySeverity(severityLevel Level) io.Writer

	// 获得一个clone的对象（线程安全）
	Clone() Logging
}

type logging struct {
	timeFormatter   func(t time.Time) string
	callerFormatter func(file string, line int, funcName string) string
	exitFunc        func(code int)
	formatter       atomic.Value
	colorFlag       int
	fileFlag        int
	fatalNoTrace    bool

	level Level

	writers sync.Map

	bufPool sync.Pool
}

var defaultLogging Logging = NewLogging()

func NewLogging(opts ...LoggingOpt) Logging {
	ret := &logging{
		timeFormatter:   TimeFormat,
		callerFormatter: CallerFormat,
		exitFunc:        os.Exit,
		//formatter:     nil,
		colorFlag:    DefaultColorFlag,
		fileFlag:     DefaultPrintFileFlag,
		fatalNoTrace: DefaultFatalNoTrace,
		level:        DefaultLevel,

		//writers: map[Level]io.Writer{},

		bufPool: sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(nil)
		}},
	}

	for k, v := range DefaultWriters {
		ret.writers.Store(k, v)
	}

	for _, v := range opts {
		v(ret)
	}
	return ret
}

func (l *logging) getCaller(depth int) string {
	if l.fileFlag != CallerNone {
		pc, file, line, ok := runtime.Caller(3 + depth)
		if !ok {
			return "???"
		}

		if (l.fileFlag & CallerShortFile) != 0 {
			file = shortFile(file)
		}

		if (l.fileFlag & CallerFileMask) == 0 {
			file = ""
			line = -1
		}
		var funcName string
		if (l.fileFlag & CallerFuncMask) != 0 {
			funcName = runtime.FuncForPC(pc).Name()
			if (l.fileFlag & CallerShortFunc) != 0 {
				idx := strings.LastIndex(funcName, ".")
				if idx != -1 && idx < (len(funcName)-1) {
					funcName = funcName[idx+1:]
				}
			}
		}
		return l.callerFormatter(file, line, funcName)
	}
	return ""
}

func (l *logging) format(writer io.Writer, level Level, depth int, keyValues KeyValues, log string) {
	caller := l.getCaller(depth)

	var (
		lvColor    string
		resetColor string
	)
	if l.colorFlag == AutoColor {
		lvColor = selectLevelColor(level)
		resetColor = ResetColor
	}

	//if log == "" || log[len(log)-1] != '\n' {
	//	log += "\n"
	//}

	formatter := l.formatter.Load()
	if formatter != nil {
		innerKvs := NewKeyValues()
		innerKvs.Add(KeyTimestamp, time.Now(), KeySeverityLevel, gLogTag[level], KeyCaller, caller)
		MergeKeyValues(innerKvs, keyValues)
		if log == "\n" {
			log = ""
		}
		innerKvs.Add(KeyContent, log)
		formatter.(Formatter).Format(writer, innerKvs)
	} else {
		writer.Write([]byte(fmt.Sprintf("%s [%s%s%s] %s %s%s",
			l.timeFormatter(time.Now()), lvColor, gLogTag[level], resetColor, caller, l.formatKeyValues(keyValues), log)))
	}
}

func (l *logging) formatKeyValues(keyValues KeyValues) string {
	if keyValues == nil || keyValues.Len() == 0 {
		return ""
	}

	buf := bytes.Buffer{}
	for _, k := range keyValues.Keys() {
		buf.WriteString(l.formatValue(keyValues.Get(k)))
		buf.WriteByte(' ')
	}
	return buf.String()
}

func (l *logging) formatValue(o interface{}) string {
	if o == nil {
		return ""
	}

	if t, ok := o.(time.Time); ok {
		if l.timeFormatter != nil {
			return l.timeFormatter(t)
		}
	}
	return formatValue(o, false)
}

func selectLevelColor(level Level) string {
	if level == INFO {
		return ForeCyan
	} else if level == WARN {
		return ForeYellow
	} else if level > WARN {
		return ForeRed
	}
	return ""
}

func (l *logging) Logf(level Level, depth int, keyValues KeyValues, format string, args ...interface{}) {
	if atomic.LoadInt32(&l.level) > level {
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
	l.format(w, level, depth, keyValues, logInfo)

	if level == PANIC {
		panic(logInfo)
	} else if level >= FATAL {
		l.processFatal(w)
	}

	//l.output(level, buf)
}

func (l *logging) Log(level Level, depth int, keyValues KeyValues, args ...interface{}) {
	if atomic.LoadInt32(&l.level) > level {
		return
	}

	logInfo := fmt.Sprint(args...)
	w := l.selectWriter(level)
	l.format(w, level, depth, keyValues, logInfo)

	if level == PANIC {
		panic(logInfo)
	} else if level >= FATAL {
		l.processFatal(w)
	}
}

func (l *logging) Logln(level Level, depth int, keyValues KeyValues, args ...interface{}) {
	if atomic.LoadInt32(&l.level) > level {
		return
	}

	logInfo := fmt.Sprintln(args...)
	w := l.selectWriter(level)
	l.format(w, level, depth, keyValues, logInfo)

	if level == PANIC {
		panic(logInfo)
	} else if level >= FATAL {
		l.processFatal(w)
	}
}

func (l *logging) getBuffer() *bytes.Buffer {
	buf := l.bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (l *logging) putBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	if buf.Len() > 256 {
		//let big buffers die a natural death.
		return
	}
	l.bufPool.Put(buf)
}

func (l *logging) processFatal(writer io.Writer) {
	if !l.fatalNoTrace {
		trace := stacks(true)
		writer.Write(trace)
	}
	l.exitFunc(-1)
}

func (l *logging) Clone() Logging {
	ret := &logging{
		timeFormatter:   l.timeFormatter,
		callerFormatter: l.callerFormatter,
		//formatter:     l.formatter,
		colorFlag:    l.colorFlag,
		fileFlag:     l.fileFlag,
		fatalNoTrace: l.fatalNoTrace,
		level:        l.level,
		//writers:       map[Level]io.Writer{},

		bufPool: sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(nil)
		}},
	}
	ret.formatter.Store(l.formatter.Load())
	l.writers.Range(func(key, value interface{}) bool {
		ret.writers.Store(key, value)
		return true
	})
	return ret
}

//func (l *logging) output(level int) {
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

func (l *logging) selectWriter(level Level) io.Writer {
	for i := level; i >= DEBUG; i-- {
		v, ok := l.writers.Load(i)
		if ok && v != nil {
			return v.(io.Writer)
		}
	}
	return os.Stdout
}

func (l *logging) SetFormatter(f Formatter) {
	l.formatter.Store(f)
}

func (l *logging) GetFormatter() Formatter {
	v := l.formatter.Load()
	if v == nil {
		return nil
	}
	return v.(Formatter)
}

func (l *logging) SetSeverityLevel(severity Level) {
	atomic.StoreInt32(&l.level, severity)
}

func (l *logging) IsEnable(severityLevel Level) bool {
	return atomic.LoadInt32(&l.level) <= severityLevel
}

// Logging不会自动为输出的Writer加锁，如果需要加锁请使用LockedWriter：
// logging.SetOutPut(&writer.LockedWriter{w})
func (l *logging) SetOutput(w io.Writer) {
	for i := DEBUG; i <= FATAL; i++ {
		l.writers.Store(i, w)
	}
}

// Logging不会自动为输出的Writer加锁，如果需要加锁请使用LockedWriter：
// logging.SetOutputBySeverity(level, &writer.LockedWriter{w})
func (l *logging) SetOutputBySeverity(severityLevel Level, w io.Writer) {
	l.writers.Store(severityLevel, w)
}

func (l *logging) GetOutputBySeverity(severityLevel Level) io.Writer {
	v, ok := l.writers.Load(severityLevel)
	if !ok {
		return nil
	}
	return v.(io.Writer)
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

func CallerFormat(file string, line int, funcName string) string {
	//no file
	if file != "" {
		if funcName == "" {
			return fmt.Sprintf("%s:%d", file, line)
		} else {
			return fmt.Sprintf("%s:%d (%s)", file, line, funcName)
		}
	} else {
		if funcName == "" {
			return ""
		} else {
			return "(" + funcName + ")"
		}
	}
}

// 配置内置Logging实现的时间格式化函数
func SetTimeFormatter(f func(t time.Time) string) func(*logging) {
	return func(logging *logging) {
		logging.timeFormatter = f
	}
}

// 配置内置Logging实现的时间格式化函数
func SetCallerFormatter(f func(file string, line int, funcName string) string) func(*logging) {
	return func(logging *logging) {
		logging.callerFormatter = f
	}
}

// 配置内置Logging Fatal退出处理函数
func SetExitFunc(f func(int)) func(*logging) {
	return func(logging *logging) {
		logging.exitFunc = f
	}
}

// 配置内置Logging实现的颜色的标志，有AutoColor、DisableColor、ForceColor
func SetColorFlag(flag int) func(*logging) {
	return func(logging *logging) {
		logging.colorFlag = flag
	}
}

// 配置内置Logging实现的文件输出标志，有ShortFile、LongFile
func SetCallerFlag(flag int) func(*logging) {
	return func(logging *logging) {
		logging.fileFlag = flag
	}
}

// 配置内置Logging实现是否在发生致命错误时打印堆栈，默认打印
func SetFatalNoTrace(noTrace bool) func(*logging) {
	return func(logging *logging) {
		logging.fatalNoTrace = noTrace
	}
}

// 设置默认Logging的日志格式化工具
func SetFormatter(f Formatter) {
	defaultLogging.SetFormatter(f)
}

// 设置默认Logging的日志严重级别
func SetSeverityLevel(severity Level) {
	defaultLogging.SetSeverityLevel(severity)
}

// 检查是否输出参数级别的日志
func IsEnable(severityLevel Level) bool {
	return defaultLogging.IsEnable(severityLevel)
}

// 设置默认Logging的输出
func SetOutput(w io.Writer) {
	defaultLogging.SetOutput(w)
}

// 设置默认Logging对应日志级别的输出
func SetOutputBySeverity(severity Level, w io.Writer) {
	defaultLogging.SetOutputBySeverity(severity, w)
}

// 获得默认Logging对应日志级别的输出
func GetOutputBySeverity(severity Level) io.Writer {
	return defaultLogging.GetOutputBySeverity(severity)
}

// 获得默认Logging
func DefaultLogging() Logging {
	return defaultLogging
}

// 使用一个Logging初始化日志系统，包括默认Logging和LoggerFactory
func Init(logging Logging) {
	defaultLogging = logging
	ResetFactoryLogging(logging)
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
