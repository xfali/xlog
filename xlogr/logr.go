// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlogr

import (
	"github.com/go-logr/logr"
	"github.com/xfali/xlog"
)

var (
	KeyLogrMessage = "LogMessage"
	KeyLogrError   = "LogError"
)

type xlogr struct {
	logging xlog.Logging
	level   xlog.Level
	field   xlog.Field
	name    string
}

func NewLogr() logr.Logger {
	return NewLogrWithLogging(xlog.DefaultLogging())
}

func NewLogrWithLogging(logging xlog.Logging) logr.Logger {
	if logging == nil {
		return nil
	}
	ret := &xlogr{
		logging: logging,
		field:   xlog.NewField(),
		level:   xlog.DefaultLevel,
	}
	//if formatter != nil {
	//	ret.logging.SetFormatter(formatter)
	//}
	return ret
}

// Enabled tests whether this Logger is enabled.  For example, commandline
// flags might be used to set the logging verbosity and disable some info
// logs.
func (l *xlogr) Enabled() bool {
	return l.logging.IsEnable(l.level)
}

// Info logs a non-error message with the given key/value pairs as context.
//
// The msg argument should be used to add some constant description to
// the log line.  The key/value pairs can then be used to add additional
// variable information.  The key/value pairs should alternate string
// keys and arbitrary values.
func (l *xlogr) Info(msg string, keysAndValues ...interface{}) {
	field := l.field.Clone()
	field.Add(KeyLogrMessage, msg)
	field.Add(keysAndValues...)
	l.logging.Logln(l.level, 1, field)
}

// Error logs an error, with the given message and key/value pairs as context.
// It functions similarly to calling Info with the "error" named value, but may
// have unique behavior, and should be preferred for logging errors (see the
// package documentations for more information).
//
// The msg field should be used to add context to any underlying error,
// while the err field should be used to attach the actual error that
// triggered this log line, if present.
func (l *xlogr) Error(err error, msg string, keysAndValues ...interface{}) {
	field := l.field.Clone()
	field.Add(KeyLogrMessage, msg, KeyLogrError, err)
	field.Add(keysAndValues...)
	l.logging.Logln(l.level, 1, field)
}

// V returns an Logger value for a specific verbosity level, relative to
// this Logger.  In other words, V values are additive.  V higher verbosity
// level means a log message is less important.  It's illegal to pass a log
// level less than zero.
func (l *xlogr) V(level int) logr.Logger {
	return &xlogr{
		level:   Int2Level(level),
		field:   l.field.Clone(),
		logging: l.logging,
		name:    l.name,
	}
}

func Int2Level(level int) xlog.Level {
	return xlog.Level(level)
}

func Level2Int(level xlog.Level) int {
	return int(level)
}

// WithValues adds some key-value pairs of context to a logger.
// See Info for documentation on how key/value pairs work.
func (l *xlogr) WithValues(keysAndValues ...interface{}) logr.Logger {
	field := l.field.Clone()
	field.Add(keysAndValues...)
	return &xlogr{
		level:   l.level,
		field:   field,
		logging: l.logging,
		name:    l.name,
	}
}

// WithName adds a new element to the logger's name.
// Successive calls with WithName continue to append
// suffixes to the logger's name.  It's strongly reccomended
// that name segments contain only letters, digits, and hyphens
// (see the package documentation for more information).
func (l *xlogr) WithName(name string) logr.Logger {
	field := l.field.Clone()
	field.Add(xlog.KeyName, name)
	ret := &xlogr{
		level:   l.level,
		logging: l.logging,
		name:    l.name + "." + name,
		field:   field,
	}

	return ret
}
