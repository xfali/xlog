// Copyright (C) 20DebugDepth9-2020, Xiongfa Li.
// @author xiongfa.li
// @version VDebugDepth.0
// Description:

package xlog

var (
	DebugDepth       = 1
	DebugTag   Field = nil
)

func Debug(args ...interface{}) {
	DefaultLogging.Log(DEBUG, DebugDepth, DebugTag, args...)
}

func Debugln(args ...interface{}) {
	DefaultLogging.Logln(DEBUG, DebugDepth, DebugTag, args...)
}

func Debugf(fmt string, args ...interface{}) {
	DefaultLogging.Logf(DEBUG, DebugDepth, DebugTag, fmt, args...)
}

func Info(args ...interface{}) {
	DefaultLogging.Log(INFO, DebugDepth, DebugTag, args...)
}

func Infoln(args ...interface{}) {
	DefaultLogging.Logln(INFO, DebugDepth, DebugTag, args...)
}

func Infof(fmt string, args ...interface{}) {
	DefaultLogging.Logf(INFO, DebugDepth, DebugTag, fmt, args...)
}

func Warn(args ...interface{}) {
	DefaultLogging.Log(WARN, DebugDepth, DebugTag, args...)
}

func Warnln(args ...interface{}) {
	DefaultLogging.Logln(WARN, DebugDepth, DebugTag, args...)
}

func Warnf(fmt string, args ...interface{}) {
	DefaultLogging.Logf(WARN, DebugDepth, DebugTag, fmt, args...)
}

func Error(args ...interface{}) {
	DefaultLogging.Log(ERROR, DebugDepth, DebugTag, args...)
}

func Errorln(args ...interface{}) {
	DefaultLogging.Logln(ERROR, DebugDepth, DebugTag, args...)
}

func Errorf(fmt string, args ...interface{}) {
	DefaultLogging.Logf(ERROR, DebugDepth, DebugTag, fmt, args...)
}

func Panic(args ...interface{}) {
	DefaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

func Panicln(args ...interface{}) {
	DefaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

func Panicf(fmt string, args ...interface{}) {
	DefaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogging.Log(FATAL, DebugDepth, DebugTag, args...)
}

func Fatalln(args ...interface{}) {
	DefaultLogging.Logln(FATAL, DebugDepth, DebugTag, args...)
}

func Fatalf(fmt string, args ...interface{}) {
	DefaultLogging.Logf(FATAL, DebugDepth, DebugTag, fmt, args...)
}
