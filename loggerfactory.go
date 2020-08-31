// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"reflect"
	"sync/atomic"
)

type LoggerFactory interface {
	GetLogger(o ...interface{}) Logger
	Reset(logging Logging) LoggerFactory
}

type loggerFactory struct {
	value atomic.Value
}

var defaultFactory = NewFactory(DefaultLogging)

func NewDefaultFactory(opts ...LoggingOpt) LoggerFactory {
	ret := &loggerFactory{}
	ret.value.Store(NewLogging(opts...))
	return ret
}

func NewFactory(logging Logging) LoggerFactory {
	ret := &loggerFactory{}
	ret.value.Store(logging)
	return ret
}

func ResetFactory(fac LoggerFactory) {
	defaultFactory = fac
}

func ResetFactoryLogging(logging Logging) {
	defaultFactory.Reset(logging)
}

func GetLogger(o ...interface{}) Logger {
	return defaultFactory.GetLogger(o...)
}

func (fac *loggerFactory) GetLogger(o ...interface{}) Logger {
	if len(o) == 0 {
		return NewLogger(fac.value.Load().(Logging))
	} else {
		t := reflect.TypeOf(o[0])
		if t.Kind() == reflect.String {
			return NewLogger(fac.value.Load().(Logging), o[0].(string))
		}
		return NewLogger(fac.value.Load().(Logging), t.String())
	}
}

func (fac *loggerFactory) Reset(logging Logging) LoggerFactory {
	fac.value.Store(logging)
	return fac
}
