// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

type LogDebug interface {
	DebugEnabled() bool
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(fmt string, args ...interface{})
}

type LogInfo interface {
	InfoEnabled() bool
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(fmt string, args ...interface{})
}

type LogWarn interface {
	WarnEnabled() bool
	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(fmt string, args ...interface{})
}

type LogError interface {
	ErrorEnabled() bool
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(fmt string, args ...interface{})
}

type LogPanic interface {
	PanicEnabled() bool
	Panic(args ...interface{})
	Panicln(args ...interface{})
	Panicf(fmt string, args ...interface{})
}

type LogFatal interface {
	FatalEnabled() bool
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

// Logger是xlog的日志封装工具，实现了常用的日志方法
type Logger interface {
	// Debug级别日志接口
	LogDebug
	// Info级别日志接口
	LogInfo
	// Warn级别日志接口
	LogWarn
	// Error级别日志接口
	LogError
	// Panic级别日志接口，注意会触发Panic
	LogPanic
	// Fatal级别日志接口，注意会触发程序退出
	LogFatal

	// 附加日志名称，注意会附加父Logger的名称，格式为：父Logger名称 + '.' + name
	WithName(name string) Logger

	// 附加日志信息，注意会附加父Logger的附加信息，如果相同则会覆盖
	WithFields(keyAndValues ...interface{}) Logger

	// 配置日志的调用深度，注意会在父Logger的基础上调整深度
	WithDepth(depth int) Logger
}
