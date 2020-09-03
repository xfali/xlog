// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"reflect"
	"strings"
	"sync/atomic"
)

type LoggerFactory interface {
	// 根据参数获得Logger
	// Param：根据默认实现，o可不填，直接返回一个没有名称的Logger。
	// 如果o有值，则只取第一个值，且当：
	// 		o为string时，使用string值作为Logger名称
	//		o为其他类型时，取package path + type name作为Logger名称，以"."分隔，如g.x.x.t.TestStructInTest
	GetLogger(o ...interface{}) Logger

	// 重置Factory的Logging（线程安全）
	Reset(logging Logging) LoggerFactory

	// 获得Factory的Logging（线程安全），可用来配置Logging（但是不建议这么做，因为Logging的配置方法不一定线程安全，更好的做法是直接重置Logging）
	// 也可以通过wrap Logging达到控制日志级别、日志输出格式的目的
	GetLogging() Logging
}

type loggerFactory struct {
	value            atomic.Value
	SimplifyNameFunc func(string) string
}

var defaultFactory LoggerFactory = NewFactory(defaultLogging)

func NewDefaultFactory(opts ...LoggingOpt) *loggerFactory {
	ret := &loggerFactory{}
	ret.value.Store(NewLogging(opts...))
	return ret
}

func NewFactory(logging Logging) *loggerFactory {
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

func GetLogging() Logging {
	return defaultFactory.GetLogging()
}

func (fac *loggerFactory) GetLogging() Logging {
	return fac.value.Load().(Logging)
}

func (fac *loggerFactory) GetLogger(o ...interface{}) Logger {
	if len(o) == 0 {
		return NewLogger(fac.value.Load().(Logging))
	} else {
		if o[0] == nil {
			return NewLogger(fac.value.Load().(Logging))
		}
		t := reflect.TypeOf(o[0])
		if t.Kind() == reflect.String {
			return NewLogger(fac.value.Load().(Logging), o[0].(string))
		}

		name := t.PkgPath()
		if name != "" {
			if fac.SimplifyNameFunc != nil {
				names := strings.Split(name, "/")
				builder := strings.Builder{}
				for _, v := range names {
					builder.WriteString(fac.SimplifyNameFunc(v))
					builder.WriteByte('.')
				}
				builder.WriteString(t.Name())
				name = builder.String()
			} else {
				name = strings.Replace(name, "/", ".", -1) + "." + t.Name()
			}
		}
		return NewLogger(fac.value.Load().(Logging), name)
	}
}

func (fac *loggerFactory) Reset(logging Logging) LoggerFactory {
	fac.value.Store(logging)
	return fac
}

func SimplifyNameFirstLetter(s string) string {
	if s == "" {
		return s
	}
	return s[:1]
}
