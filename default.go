// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

type xlog struct {
	logging Logging
	depth   int
	field   Field
	name    string
}

func NewLogger(logging Logging, name ...string) Logger {
	return newLogger(logging, nil, name...)
}

func newLogger(logging Logging, field Field, name ...string) *xlog {
	if field == nil {
		field = NewField()
	}
	var t string
	if len(name) > 0 {
		t = name[0]
		if t != "" {
			field.Add(KeyName, t)
		}
	}
	return &xlog{
		logging: logging,
		depth:   1,
		name:    t,
		field:   field,
	}
}

func (l *xlog) Debug(args ...interface{}) {
	l.logging.Log(DEBUG, l.depth, l.field, args...)
}

func (l *xlog) Debugln(args ...interface{}) {
	l.logging.Logln(DEBUG, l.depth, l.field, args...)
}

func (l *xlog) Debugf(fmt string, args ...interface{}) {
	l.logging.Logf(DEBUG, l.depth, l.field, fmt, args...)
}

func (l *xlog) Info(args ...interface{}) {
	l.logging.Log(INFO, l.depth, l.field, args...)
}

func (l *xlog) Infoln(args ...interface{}) {
	l.logging.Logln(INFO, l.depth, l.field, args...)
}

func (l *xlog) Infof(fmt string, args ...interface{}) {
	l.logging.Logf(INFO, l.depth, l.field, fmt, args...)
}

func (l *xlog) Warn(args ...interface{}) {
	l.logging.Log(WARN, l.depth, l.field, args...)
}

func (l *xlog) Warnln(args ...interface{}) {
	l.logging.Logln(WARN, l.depth, l.field, args...)
}

func (l *xlog) Warnf(fmt string, args ...interface{}) {
	l.logging.Logf(WARN, l.depth, l.field, fmt, args...)
}

func (l *xlog) Error(args ...interface{}) {
	l.logging.Log(ERROR, l.depth, l.field, args...)
}

func (l *xlog) Errorln(args ...interface{}) {
	l.logging.Logln(ERROR, l.depth, l.field, args...)
}

func (l *xlog) Errorf(fmt string, args ...interface{}) {
	l.logging.Logf(ERROR, l.depth, l.field, fmt, args...)
}

func (l *xlog) Panic(args ...interface{}) {
	l.logging.Log(PANIC, l.depth, l.field, args...)
}

func (l *xlog) Panicln(args ...interface{}) {
	l.logging.Log(PANIC, l.depth, l.field, args...)
}

func (l *xlog) Panicf(fmt string, args ...interface{}) {
	l.logging.Log(PANIC, l.depth, l.field, args...)
}

func (l *xlog) Fatal(args ...interface{}) {
	l.logging.Log(FATAL, l.depth, l.field, args...)
}

func (l *xlog) Fatalln(args ...interface{}) {
	l.logging.Logln(FATAL, l.depth, l.field, args...)
}

func (l *xlog) Fatalf(fmt string, args ...interface{}) {
	l.logging.Logf(FATAL, l.depth, l.field, fmt, args...)
}

func (l *xlog) WithName(name string) Logger {
	if l == nil {
		return nil
	}

	if l.name != "" {
		name = l.name + "." + name
	}
	ret := newLogger(l.logging, l.field.Clone(), name)
	ret.field.Add(KeyName, ret.name)
	ret.depth = l.depth

	return ret
}

func (l *xlog) WithFields(keyAndValues ...interface{}) Logger {
	if l == nil {
		return nil
	}
	ret := newLogger(l.logging, l.field.Clone(), l.name)
	ret.field.Add(keyAndValues...)
	ret.depth = l.depth

	return ret
}

func (l *xlog) WithDepth(depth int) Logger {
	if l == nil {
		return nil
	}
	ret := newLogger(l.logging, l.field.Clone(), l.name)
	ret.depth += depth

	return ret
}
