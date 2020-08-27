// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"reflect"
	"sync/atomic"
)

type LoggerFactory struct {
	value atomic.Value
}

var defaultFactory = newFactory(defaultLogging)

func NewFactory(opts ...LoggingOpt) *LoggerFactory {
	ret := &LoggerFactory{}
	ret.value.Store(NewLogging(opts...))
	return ret
}

func newFactory(logging *Logging) *LoggerFactory {
	ret := &LoggerFactory{}
	ret.value.Store(logging)
	return ret
}

func ResetFactory(logging *Logging) {
	defaultFactory.Reset(logging)
}

func GetLogger(o ...interface{}) Logger {
	return defaultFactory.GetLogger(o...)
}

func (fac *LoggerFactory) GetLogger(o ...interface{}) Logger {
	if len(o) == 0 {
		return NewLogger(fac.value.Load().(*Logging))
	} else {
		t := reflect.TypeOf(o[0])
		if t.Kind() == reflect.String {
			return NewLogger(fac.value.Load().(*Logging), o[0].(string))
		}
		return NewLogger(fac.value.Load().(*Logging), t.String())
	}
}

func (fac *LoggerFactory) Reset(logging *Logging) *LoggerFactory {
	fac.value.Store(logging)
	return fac
}
