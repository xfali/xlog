// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import "io"

type LevelHook func(Level) Level

type hookLevelLogging struct {
	logging Logging
	hook    LevelHook
}

func NewHookLevelLogging(logging Logging, hook LevelHook) *hookLevelLogging {
	return &hookLevelLogging{
		logging: logging,
		hook:    hook,
	}
}

func (l *hookLevelLogging) Logf(level Level, depth int, field Field, format string, args ...interface{}) {
	l.logging.Logf(l.hook(level), depth+1, field, format, args...)
}

func (l *hookLevelLogging) Log(level Level, depth int, field Field, args ...interface{}) {
	l.logging.Log(l.hook(level), depth+1, field, args...)
}

func (l *hookLevelLogging) Logln(level Level, depth int, field Field, args ...interface{}) {
	l.logging.Logln(l.hook(level), depth+1, field, args...)
}

func (l *hookLevelLogging) SetFormatter(f Formatter) {
	l.logging.SetFormatter(f)
}

func (l *hookLevelLogging) SetSeverityLevel(severityLevel Level) {
	l.logging.SetSeverityLevel(severityLevel)
}

func (l *hookLevelLogging) IsEnable(severityLevel Level) bool {
	return l.logging.IsEnable(severityLevel)
}

func (l *hookLevelLogging) SetOutput(w io.Writer) {
	l.logging.SetOutput(w)
}

func (l *hookLevelLogging) SetOutputBySeverity(severityLevel Level, w io.Writer) {
	l.logging.SetOutputBySeverity(severityLevel, w)
}
