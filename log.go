// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

type Debug interface {
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(fmt string, args ...interface{})
}

type Info interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(fmt string, args ...interface{})
}

type Warn interface {
	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(fmt string, args ...interface{})
}

type Error interface {
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(fmt string, args ...interface{})
}

type Fatal interface {
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

type Logger interface {
	Debug
	Info
	Warn
	Error
	Fatal
}
