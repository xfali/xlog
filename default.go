// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

type xlog struct {
	logging *Logging
	depth   int
}

func NewLogger(logging *Logging) Logger {
	return newLogger(logging)
}

func newLogger(logging *Logging) *xlog {
	return &xlog{
		logging: logging,
		depth:   1,
	}
}

func (l *xlog) Debug(args ...interface{}) {
	l.logging.Log(DEBUG, l.depth, args...)
}

func (l *xlog) Debugln(args ...interface{}) {
	l.logging.Logln(DEBUG, l.depth, args...)
}

func (l *xlog) Debugf(fmt string, args ...interface{}) {
	l.logging.Logf(DEBUG, l.depth, fmt, args...)
}

func (l *xlog) Info(args ...interface{}) {
	l.logging.Log(INFO, l.depth, args...)
}

func (l *xlog) Infoln(args ...interface{}) {
	l.logging.Logln(INFO, l.depth, args...)
}

func (l *xlog) Infof(fmt string, args ...interface{}) {
	l.logging.Logf(INFO, l.depth, fmt, args...)
}

func (l *xlog) Warn(args ...interface{}) {
	l.logging.Log(WARN, l.depth, args...)
}

func (l *xlog) Warnln(args ...interface{}) {
	l.logging.Logln(WARN, l.depth, args...)
}

func (l *xlog) Warnf(fmt string, args ...interface{}) {
	l.logging.Logf(WARN, l.depth, fmt, args...)
}

func (l *xlog) Error(args ...interface{}) {
	l.logging.Log(ERROR, l.depth, args...)
}

func (l *xlog) Errorln(args ...interface{}) {
	l.logging.Logln(ERROR, l.depth, args...)
}

func (l *xlog) Errorf(fmt string, args ...interface{}) {
	l.logging.Logf(ERROR, l.depth, fmt, args...)
}

func (l *xlog) Fatal(args ...interface{}) {
	l.logging.Log(FATAL, l.depth, args...)
}

func (l *xlog) Fatalln(args ...interface{}) {
	l.logging.Logln(FATAL, l.depth, args...)
}

func (l *xlog) Fatalf(fmt string, args ...interface{}) {
	l.logging.Logf(FATAL, l.depth, fmt, args...)
}

func (l *xlog) NewDepth(depth int) Logger {
	if l == nil {
		return nil
	}
	ret := newLogger(l.logging)
	ret.depth += depth

	return ret
}
