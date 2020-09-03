// Copyright (C) 20DebugDepth9-2020, Xiongfa Li.
// @author xiongfa.li
// @version VDebugDepth.0
// Description:

package xlog

var (
	// 默认日志深度
	DebugDepth = 1
	// 默认日志附加信息（无）
	DebugTag Field = nil
)

// 使用默认的Logging，输出Debug级别的日志
func Debug(args ...interface{}) {
	defaultLogging.Log(DEBUG, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Debug级别的日志
func Debugln(args ...interface{}) {
	defaultLogging.Logln(DEBUG, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Debug级别的日志
func Debugf(fmt string, args ...interface{}) {
	defaultLogging.Logf(DEBUG, DebugDepth, DebugTag, fmt, args...)
}

// 使用默认的Logging，输出Info级别的日志
func Info(args ...interface{}) {
	defaultLogging.Log(INFO, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Info级别的日志
func Infoln(args ...interface{}) {
	defaultLogging.Logln(INFO, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Info级别的日志
func Infof(fmt string, args ...interface{}) {
	defaultLogging.Logf(INFO, DebugDepth, DebugTag, fmt, args...)
}

// 使用默认的Logging，输出Warn级别的日志
func Warn(args ...interface{}) {
	defaultLogging.Log(WARN, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Warn级别的日志
func Warnln(args ...interface{}) {
	defaultLogging.Logln(WARN, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Warn级别的日志
func Warnf(fmt string, args ...interface{}) {
	defaultLogging.Logf(WARN, DebugDepth, DebugTag, fmt, args...)
}

// 使用默认的Logging，输出Error级别的日志
func Error(args ...interface{}) {
	defaultLogging.Log(ERROR, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Error级别的日志
func Errorln(args ...interface{}) {
	defaultLogging.Logln(ERROR, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Error级别的日志
func Errorf(fmt string, args ...interface{}) {
	defaultLogging.Logf(ERROR, DebugDepth, DebugTag, fmt, args...)
}

// 使用默认的Logging，输出Panic级别的日志，注意会触发panic
func Panic(args ...interface{}) {
	defaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Panic级别的日志，注意会触发panic
func Panicln(args ...interface{}) {
	defaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Panic级别的日志，注意会触发panic
func Panicf(fmt string, args ...interface{}) {
	defaultLogging.Log(PANIC, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Fatal级别的日志，注意会触发程序退出
func Fatal(args ...interface{}) {
	defaultLogging.Log(FATAL, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Fatal级别的日志，注意会触发程序退出
func Fatalln(args ...interface{}) {
	defaultLogging.Logln(FATAL, DebugDepth, DebugTag, args...)
}

// 使用默认的Logging，输出Fatal级别的日志，注意会触发程序退出
func Fatalf(fmt string, args ...interface{}) {
	defaultLogging.Logf(FATAL, DebugDepth, DebugTag, fmt, args...)
}

// 配置默认的调用深度
func WithDepth(depth int) {
	DebugDepth = depth
}

// 配置默认附件信息
func WithFields(keyAndValues ...interface{}) {
	DebugTag = NewField(keyAndValues...)
}
