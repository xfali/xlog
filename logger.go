// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

type LogDebug interface {
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(fmt string, args ...interface{})
}

type LogInfo interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(fmt string, args ...interface{})
}

type LogWarn interface {
	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(fmt string, args ...interface{})
}

type LogError interface {
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(fmt string, args ...interface{})
}

type LogFatal interface {
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

type Logger interface {
	LogDebug
	LogInfo
	LogWarn
	LogError
	LogFatal

	WithName(name string) Logger
	WithDepth(depth int) Logger
}
