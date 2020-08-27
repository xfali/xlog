// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import "reflect"

type LoggerFactory struct {
	logging *Logging
}

func (fac *LoggerFactory) GetLogger(o interface{}) Logger {
	if o == nil {
		return NewLogger(fac.logging)
	}
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	//path := t.PkgPath()
	return nil
}
