// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"github.com/xfali/xlog/value"
	"reflect"
	"strings"
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

	// 获得Factory的Logging（线程安全），可用来配置Logging
	// 也可以通过wrap Logging达到控制日志级别、日志输出格式的目的
	GetLogging() Logging
}

type loggerFactory struct {
	value            value.Value
	SimplifyNameFunc func(string) string
}

var defaultFactory value.Value = value.NewSimpleValue(LoggerFactory(NewMutableFactory(DefaultLogging())))

func NewDefaultFactory(opts ...LoggingOpt) *loggerFactory {
	return NewFactory(NewLogging(opts...))
}

func NewFactory(logging Logging) *loggerFactory {
	return NewFactoryWithValue(value.NewAtomicValue(logging))
}

func NewFactoryWithValue(v value.Value) *loggerFactory {
	ret := &loggerFactory{
		value: v,
	}
	return ret
}

// 重新配置全局的默认LoggerFactory，该方法同时会重置全局的默认Logging
// 由于线程安全性受defaultLogging、defaultFactory初始化（调用InitOnce）的Value决定，
// 所以需要确定是否确实需要调用该方法重置Logging，并保证Value线程安全
func ResetFactory(fac LoggerFactory) {
	defaultFactory.Store(fac)
	ResetLogging(fac.GetLogging())
}

// 重新配置全局的默认Logging，该方法同时会重置全局的默认LoggerFactory的Logging
// 由于线程安全性受defaultLogging、defaultFactory初始化（调用InitOnce）的Value决定，
// 所以需要确定是否确实需要调用该方法重置Logging，并保证Value线程安全
func ResetLogging(logging Logging) {
	defaultLogging.Store(logging)
	defaultFactory.Load().(LoggerFactory).Reset(defaultLogging.Load().(Logging))
}

// 通过全局默认LoggerFactory获取Logger
// Param：根据默认实现，o可不填，直接返回一个没有名称的Logger。
// 如果o有值，则只取第一个值，且当：
// 		o为string时，使用string值作为Logger名称
//		o为其他类型时，取package path + type name作为Logger名称，以"."分隔，如g.x.x.t.TestStructInTest
func GetLogger(o ...interface{}) Logger {
	return defaultFactory.Load().(LoggerFactory).GetLogger(o...)
}

// 通过全局默认LoggerFactory的Logging
func GetLogging() Logging {
	return defaultFactory.Load().(LoggerFactory).GetLogging()
}

func (fac *loggerFactory) GetLogging() Logging {
	return fac.value.Load().(Logging)
}

func (fac *loggerFactory) GetLogger(o ...interface{}) Logger {
	name := getObjectName(fac.SimplifyNameFunc, o...)
	return newLogger(fac.value.Load().(Logging), nil, name)
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

type mutableLoggerFactory struct {
	loggerFactory
}

func NewMutableFactory(logging Logging) *mutableLoggerFactory {
	return NewMutableFactoryWithValue(value.NewAtomicValue(logging))
}

func NewMutableFactoryWithValue(v value.Value) *mutableLoggerFactory {
	ret := &mutableLoggerFactory{}
	ret.value = v
	return ret
}

func (fac *mutableLoggerFactory) GetLogger(o ...interface{}) Logger {
	name := getObjectName(fac.SimplifyNameFunc, o...)
	return newMutableLogger(fac.value, nil, name)
}

func getObjectName(simpleFunc func(string) string, o ...interface{}) string {
	if len(o) == 0 {
		return ""
	} else {
		if o[0] == nil {
			return ""
		}
		t := reflect.TypeOf(o[0])
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() == reflect.String {
			return o[0].(string)
		}

		name := t.PkgPath()
		if name != "" {
			if simpleFunc != nil {
				names := strings.Split(name, "/")
				builder := strings.Builder{}
				for _, v := range names {
					builder.WriteString(simpleFunc(v))
					builder.WriteByte('.')
				}
				builder.WriteString(t.Name())
				name = builder.String()
			} else {
				name = strings.Replace(name, "/", ".", -1) + "." + t.Name()
			}
		}
		return name
	}
}
