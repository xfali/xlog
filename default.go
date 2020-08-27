// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

type xlog struct {
	logging *Logging
	depth   int
	tag     string
}

func NewLogger(logging *Logging, tag ...string) Logger {
	return newLogger(logging, tag...)
}

func newLogger(logging *Logging, tag ...string) *xlog {
	var t string
	if len(tag) > 0 {
		t = tag[0]
	}
	return &xlog{
		logging: logging,
		depth:   1,
		tag:     t,
	}
}

func (l *xlog) Debug(args ...interface{}) {
	l.logging.Log(DEBUG, l.depth, l.tag, args...)
}

func (l *xlog) Debugln(args ...interface{}) {
	l.logging.Logln(DEBUG, l.depth, l.tag, args...)
}

func (l *xlog) Debugf(fmt string, args ...interface{}) {
	l.logging.Logf(DEBUG, l.depth, l.tag, fmt, args...)
}

func (l *xlog) Info(args ...interface{}) {
	l.logging.Log(INFO, l.depth, l.tag, args...)
}

func (l *xlog) Infoln(args ...interface{}) {
	l.logging.Logln(INFO, l.depth, l.tag, args...)
}

func (l *xlog) Infof(fmt string, args ...interface{}) {
	l.logging.Logf(INFO, l.depth, l.tag, fmt, args...)
}

func (l *xlog) Warn(args ...interface{}) {
	l.logging.Log(WARN, l.depth, l.tag, args...)
}

func (l *xlog) Warnln(args ...interface{}) {
	l.logging.Logln(WARN, l.depth, l.tag, args...)
}

func (l *xlog) Warnf(fmt string, args ...interface{}) {
	l.logging.Logf(WARN, l.depth, l.tag, fmt, args...)
}

func (l *xlog) Error(args ...interface{}) {
	l.logging.Log(ERROR, l.depth, l.tag, args...)
}

func (l *xlog) Errorln(args ...interface{}) {
	l.logging.Logln(ERROR, l.depth, l.tag, args...)
}

func (l *xlog) Errorf(fmt string, args ...interface{}) {
	l.logging.Logf(ERROR, l.depth, l.tag, fmt, args...)
}

func (l *xlog) Fatal(args ...interface{}) {
	l.logging.Log(FATAL, l.depth, l.tag, args...)
}

func (l *xlog) Fatalln(args ...interface{}) {
	l.logging.Logln(FATAL, l.depth, l.tag, args...)
}

func (l *xlog) Fatalf(fmt string, args ...interface{}) {
	l.logging.Logf(FATAL, l.depth, l.tag, fmt, args...)
}

func (l *xlog) WithName(name string) Logger {
	if l == nil {
		return nil
	}
	ret := newLogger(l.logging, name)
	ret.depth = l.depth

	return ret
}

func (l *xlog) WithDepth(depth int) Logger {
	if l == nil {
		return nil
	}
	ret := newLogger(l.logging, l.tag)
	ret.depth += depth

	return ret
}

