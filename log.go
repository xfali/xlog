// Copyright (C) 20DebugDepth9-2020, Xiongfa Li.
// @author xiongfa.li
// @version VDebugDepth.0
// Description:

package xlog

import (
	"github.com/xfali/xlog/value"
)

var (
	// 默认日志深度
	LogDepth value.Value = value.NewSimpleValue(int(1))
	// 默认日志附加信息（无）
	LogFields value.Value = value.NewSimpleValue(nil)
)

// 使用默认的Logging，输出Debug级别的日志
func Debug(args ...interface{}) {
	DefaultLogging().Log(DEBUG, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Debug级别的日志
func Debugln(args ...interface{}) {
	DefaultLogging().Logln(DEBUG, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Debug级别的日志
func Debugf(fmt string, args ...interface{}) {
	DefaultLogging().Logf(DEBUG, LogDepth.Load().(int), getLogField(), fmt, args...)
}

// 使用默认的Logging，输出Info级别的日志
func Info(args ...interface{}) {
	DefaultLogging().Log(INFO, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Info级别的日志
func Infoln(args ...interface{}) {
	DefaultLogging().Logln(INFO, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Info级别的日志
func Infof(fmt string, args ...interface{}) {
	DefaultLogging().Logf(INFO, LogDepth.Load().(int), getLogField(), fmt, args...)
}

// 使用默认的Logging，输出Warn级别的日志
func Warn(args ...interface{}) {
	DefaultLogging().Log(WARN, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Warn级别的日志
func Warnln(args ...interface{}) {
	DefaultLogging().Logln(WARN, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Warn级别的日志
func Warnf(fmt string, args ...interface{}) {
	DefaultLogging().Logf(WARN, LogDepth.Load().(int), getLogField(), fmt, args...)
}

// 使用默认的Logging，输出Error级别的日志
func Error(args ...interface{}) {
	DefaultLogging().Log(ERROR, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Error级别的日志
func Errorln(args ...interface{}) {
	DefaultLogging().Logln(ERROR, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Error级别的日志
func Errorf(fmt string, args ...interface{}) {
	DefaultLogging().Logf(ERROR, LogDepth.Load().(int), getLogField(), fmt, args...)
}

// 使用默认的Logging，输出Panic级别的日志，注意会触发panic
func Panic(args ...interface{}) {
	DefaultLogging().Log(PANIC, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Panic级别的日志，注意会触发panic
func Panicln(args ...interface{}) {
	DefaultLogging().Logln(PANIC, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Panic级别的日志，注意会触发panic
func Panicf(fmt string, args ...interface{}) {
	DefaultLogging().Logf(PANIC, LogDepth.Load().(int), getLogField(), fmt, args...)
}

// 使用默认的Logging，输出Fatal级别的日志，注意会触发程序退出
func Fatal(args ...interface{}) {
	DefaultLogging().Log(FATAL, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Fatal级别的日志，注意会触发程序退出
func Fatalln(args ...interface{}) {
	DefaultLogging().Logln(FATAL, LogDepth.Load().(int), getLogField(), args...)
}

// 使用默认的Logging，输出Fatal级别的日志，注意会触发程序退出
func Fatalf(fmt string, args ...interface{}) {
	DefaultLogging().Logf(FATAL, LogDepth.Load().(int), getLogField(), fmt, args...)
}

func getLogField() KeyValues {
	ret := LogFields.Load()
	if ret == nil {
		return nil
	}
	return ret.(KeyValues)
}

// 配置默认的调用深度
func WithDepth(depth int) {
	LogDepth.Store(depth)
}

// 配置默认附件信息
func WithFields(keyAndValues ...interface{}) {
	LogFields.Store(NewKeyValues(keyAndValues...))
}
