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
	"io"
	"time"
)

type Iterator interface {
	HasNext() bool
	Next() (string, interface{})
}

type KeyValues interface {
	Add(keyAndValues ...interface{}) error
	GetAll() map[string]interface{}
	Keys() []string
	Get(key string) interface{}
	Remove(key string) error
	Len() int

	Iterator() Iterator

	Clone() KeyValues
}

type Formatter interface {
	//将日志keyValues格式化
	Format(writer io.Writer, keyValues KeyValues) error
}

type defaultKeyValues [2]interface{}

func NewKeyValues(keyAndValues ...interface{}) KeyValues {
	ret := &defaultKeyValues{}
	ret[0] = []string{}
	ret[1] = map[string]interface{}{}

	ret.Add(keyAndValues...)
	return ret
}

func (f *defaultKeyValues) Add(keyAndValues ...interface{}) error {
	size := len(keyAndValues)
	if size == 0 {
		return nil
	}
	//keys := f[0].([]string)
	kvs := f[1].(map[string]interface{})
	var k string
	for i := 0; i < size; i++ {
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
				f[0] = append(f[0].([]string), k)
			}
		} else {
			kvs[k] = keyAndValues[i]
		}
	}
	return nil
}

func (f defaultKeyValues) GetAll() map[string]interface{} {
	return f[1].(map[string]interface{})
}

func (f defaultKeyValues) Iterator() Iterator {
	return &defaultIterator{
		keyValues: f,
		cur:       0,
	}
}

func (f defaultKeyValues) Keys() []string {
	return f[0].([]string)
}

func (f defaultKeyValues) Get(key string) interface{} {
	return f[1].(map[string]interface{})[key]
}

func (f *defaultKeyValues) Remove(key string) error {
	_, ok := f[1].(map[string]interface{})[key]
	if ok {
		delete(f[1].(map[string]interface{}), key)
		keys := f[0].([]string)
		for i := 0; i < len(keys); i++ {
			if keys[i] == key {
				f[0] = append(keys[:i], keys[i+1:]...)
				break
			}
		}
		return nil
	} else {
		return errors.New("Key not found ")
	}
}

func (f defaultKeyValues) Len() int {
	return len(f[0].([]string))
}

func (f defaultKeyValues) Clone() KeyValues {
	ret := &defaultKeyValues{}
	keys := make([]string, len(f[0].([]string)))
	kvs := make(map[string]interface{}, len(f[1].(map[string]interface{})))

	copy(keys, f[0].([]string))
	for _, k := range keys {
		kvs[k] = f[1].(map[string]interface{})[k]
	}
	ret[0] = keys
	ret[1] = kvs

	return ret
}

type defaultIterator struct {
	keyValues defaultKeyValues
	cur       int
}

func (c *defaultIterator) HasNext() bool {
	return c.cur < len(c.keyValues[0].([]string))
}

func (c *defaultIterator) Next() (string, interface{}) {
	v := c.keyValues[0].([]string)[c.cur]
	c.cur++
	return v, c.keyValues[1].(map[string]interface{})[v]
}

func MergeKeyValues(keyValues ...KeyValues) (KeyValues, error) {
	if len(keyValues) == 0 {
		return nil, errors.New("No keyValues to merge ")
	} else {
		kvs := keyValues[0]
		var tmp KeyValues
		for i := 1; i < len(keyValues); i++ {
			tmp = keyValues[i]
			if tmp == nil {
				continue
			}
			keys := tmp.Keys()
			for _, k := range keys {
				err := kvs.Add(k, tmp.Get(k))
				if err != nil {
					return kvs, err
				}
			}
		}
		return kvs, nil
	}
}

type TextFormatter struct {
	TimeFormat func(t time.Time) string
	WithQuote  bool
	SortFunc   func([]string)
}

func (f *TextFormatter) Format(writer io.Writer, keyValues KeyValues) error {
	keys := keyValues.Keys()
	if len(keys) == 0 {
		return nil
	}

	if f.SortFunc != nil {
		f.SortFunc(keys)
	}

	buf := bytes.Buffer{}
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(f.formatValue(keyValues.Get(k)))
		buf.WriteByte(' ')
	}
	if buf.Cap() == 0 || buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}
	_, err := writer.Write(buf.Bytes())
	return err
}

func (f *TextFormatter) formatValue(o interface{}) string {
	if o == nil {
		return ""
	}

	if t, ok := o.(time.Time); ok {
		if f.TimeFormat != nil {
			o = f.TimeFormat(t)
		}
	}
	return formatValue(o, f.WithQuote)
}

func formatValue(o interface{}, quote bool) string {
	if o == nil {
		return ""
	}

	var ret string
	if s, ok := o.(string); ok {
		ret = s
	} else {
		ret = fmt.Sprint(o)
	}

	if quote {
		ret = fmt.Sprintf("%q", ret)
	}
	return ret
}

type JsonFormatter struct {
}

func (f *JsonFormatter) Format(writer io.Writer, keyValues KeyValues) error {
	d, err := json.Marshal(keyValues.GetAll())
	if err != nil {
		return err
	}
	_, err = writer.Write(d)
	return err
}
