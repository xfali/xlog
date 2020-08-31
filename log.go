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
	defaultLogging.Log(DEBUG, DebugDepth, DebugTag, args...)
}

func Debugln(args ...interface{}) {
	defaultLogging.Logln(DEBUG, DebugDepth, DebugTag, args...)
}

func Debugf(fmt string, args ...interface{}) {
	defaultLogging.Logf(DEBUG, DebugDepth, DebugTag, fmt, args...)
}

func Info(args ...interface{}) {
	defaultLogging.Log(INFO, DebugDepth, DebugTag, args...)
}

func Infoln(args ...interface{}) {
	defaultLogging.Logln(INFO, DebugDepth, DebugTag, args...)
}

func Infof(fmt string, args ...interface{}) {
	defaultLogging.Logf(INFO, DebugDepth, DebugTag, fmt, args...)
}

func Warn(args ...interface{}) {
	defaultLogging.Log(WARN, DebugDepth, DebugTag, args...)
}

func Warnln(args ...interface{}) {
	defaultLogging.Logln(WARN, DebugDepth, DebugTag, args...)
}

func Warnf(fmt string, args ...interface{}) {
	defaultLogging.Logf(WARN, DebugDepth, DebugTag, fmt, args...)
}

func Error(args ...interface{}) {
	defaultLogging.Log(ERROR, DebugDepth, DebugTag, args...)
}

func Errorln(args ...interface{}) {
	defaultLogging.Logln(ERROR, DebugDepth, DebugTag, args...)
}

func Errorf(fmt string, args ...interface{}) {
	defaultLogging.Logf(ERROR, DebugDepth, DebugTag, fmt, args...)
}

func Panic(args ...interface{}) {
	defaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

func Panicln(args ...interface{}) {
	defaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

func Panicf(fmt string, args ...interface{}) {
	defaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

func Fatal(args ...interface{}) {
	defaultLogging.Log(FATAL, DebugDepth, DebugTag, args...)
}

func Fatalln(args ...interface{}) {
	defaultLogging.Logln(FATAL, DebugDepth, DebugTag, args...)
}

func Fatalf(fmt string, args ...interface{}) {
	defaultLogging.Logf(FATAL, DebugDepth, DebugTag, fmt, args...)
}
