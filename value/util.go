// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package value

import "sync/atomic"

type Value interface {
	Set(interface{})
	Get() interface{}
}

type SimpleValue struct {
	o interface{}
}

func NewSimpleValue(o interface{}) *SimpleValue {
	return &SimpleValue{o: o}
}

func (l *SimpleValue) Get() interface{} {
	return l.o
}

func (l *SimpleValue) Set(o interface{}) {
	l.o = o
}

type AtomicValue atomic.Value

func NewAtomicValue(o interface{}) *AtomicValue {
	ret := &AtomicValue{}
	if o != nil {
		ret.Set(o)
	}
	return ret
}
func (l *AtomicValue) Get() interface{} {
	return (*atomic.Value)(l).Load()
}

func (l *AtomicValue) Set(o interface{}) {
	(*atomic.Value)(l).Store(o)
}
