// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Iterator interface {
	HasNext() bool
	Next() (string, interface{})
}

type Filed interface {
	Add(keyAndValues ...interface{}) error
	GetAll() map[string]interface{}
	Keys() []string
	Get(key string) interface{}
	Iterator() Iterator
}

type Formatter interface {
	Format(Filed) ([]byte, error)
}

type defaultField [2]interface{}

func NewField() Filed {
	ret := &defaultField{
	}
	ret[0] = []string{}
	ret[1] = map[string]interface{}{}
	return ret
}

func (f *defaultField) Add(keyAndValues ...interface{}) error {
	size := len(keyAndValues)
	if size == 0 {
		return nil
	}
	//keys := f[0].([]string)
	kvs := f[1].(map[string]interface{})
	var k string
	for i := 0; i < size; i ++ {
		if i%2 == 0 {
			if keyAndValues[i] == nil {
				return errors.New("Key must be not nil ")
			}
			s, ok := keyAndValues[i].(string)
			if !ok {
				return errors.New("Key must be string ")
			}
			k = s
			if _, ok := kvs[k]; !ok {
				f[0]= append(f[0].([]string), k)
			}
		} else {
			kvs[k] = keyAndValues[i]
		}
	}
	return nil
}

func (f *defaultField) GetAll() map[string]interface{} {
	return f[1].(map[string]interface{})
}

func (f *defaultField) Iterator() Iterator {
	return &defaultIterator{
		field: f,
		cur:   0,
	}
}

func (f defaultField) Keys() []string {
	return f[0].([]string)
}

func (f defaultField) Get(key string) interface{} {
	return f[1].(map[string]interface{})[key]
}

type defaultIterator struct {
	field *defaultField
	cur   int
}

func (c *defaultIterator) HasNext() bool {
	if c.cur < len(c.field[0].([]string)) {
		return true
	}
	return false
}

func (c *defaultIterator) Next() (string, interface{}) {
	v := c.field[0].([]string)[c.cur]
	c.cur++
	return v, c.field[1].(map[string]interface{})[v]
}

type BaseFormatter struct {
}

type TextFormatter struct {
	SortFunc func([]string)
}

func (f *TextFormatter) Format(filed Filed) ([]byte, error) {
	keys := filed.Keys()
	if len(keys) == 0 {
		return nil, nil
	}

	if f.SortFunc != nil {
		f.SortFunc(keys)
	}

	buf := bytes.Buffer{}
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(formatValue(filed.Get(k)))
		buf.WriteByte(' ')
	}

	return buf.Bytes(), nil
}

func formatValue(o interface{}) string {
	if o == nil {
		return ""
	}

	if s, ok := o.(string); ok {
		return s
	}
	return fmt.Sprint(o)
}

type JsonFormatter struct {
}

func (f *JsonFormatter) Format(filed Filed) ([]byte, error) {
	return json.Marshal(filed.GetAll())
}
