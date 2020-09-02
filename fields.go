// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"errors"
	"io"
)

type Field []interface{}

func NewField(keyAndValues ...interface{}) Field {
	ret := make(Field, 0, len(keyAndValues))
	ret = append(ret, keyAndValues...)
	return ret
}

func NewFieldWithCap(size int) Field {
	ret := make(Field, 0, size)
	return ret
}

func ToKeyValues(field Field) defaultField {
	d := defaultField{}
	d[0] = []string{}
	d[1] = map[string]interface{}{}

	d.Add(field...)
	return d
}

func (field Field) Add(keyAndValues ...interface{}) Field {
	return append(field, keyAndValues...)
}

func (field Field) Clone() Field {
	ret := make(Field, len(field))
	copy(ret, field)
	return ret
}

func (field Field) Keys() []string {
	ret := make([]string, 0, len(field)+1/2)
	for i := 0; i < len(field); i += 2 {
		ret = append(ret, field[i].(string))
	}
	return ret
}

func (field Field) Values() []interface{} {
	ret := make([]interface{}, 0, len(field)+1/2)
	for i := 1; i < len(field); i += 2 {
		ret = append(ret, field[i])
	}
	return ret
}

type Iterator interface {
	HasNext() bool
	Next() (string, interface{})
}

type KeyAndValues interface {
	Add(keyAndValues ...interface{}) error
	GetAll() map[string]interface{}
	Keys() []string
	Get(key string) interface{}
	Iterator() Iterator

	Clone() KeyAndValues
}

type Formatter interface {
	//将日志Field格式化
	Format(writer io.Writer, field Field) error
}

type defaultField [2]interface{}

func NewKeyValues(keyAndValues ...interface{}) KeyAndValues {
	ret := &defaultField{}
	ret[0] = []string{}
	ret[1] = map[string]interface{}{}

	ret.Add(keyAndValues...)
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

func (f defaultField) GetAll() map[string]interface{} {
	return f[1].(map[string]interface{})
}

func (f defaultField) Iterator() Iterator {
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

func (f defaultField) Clone() KeyAndValues {
	ret := NewKeyValues()

	for _, k := range f.Keys() {
		ret.Add(k, f.Get(k))
	}

	return ret
}

type defaultIterator struct {
	field defaultField
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

func MergeFields(fields ...Field) Field {
	if len(fields) == 0 {
		return nil
	} else {
		field := fields[0]
		var tmp Field
		for i := 1; i < len(fields); i++ {
			tmp = fields[i]
			if tmp == nil {
				continue
			}
			field = append(field, tmp...)
		}
		return field
	}
}

func MergeKeyValues(kvs ...KeyAndValues) (KeyAndValues, error) {
	if len(kvs) == 0 {
		return nil, errors.New("No field to merge ")
	} else {
		field := kvs[0]
		var tmp KeyAndValues
		for i := 1; i < len(kvs); i++ {
			tmp = kvs[i]
			if tmp == nil {
				continue
			}
			keys := tmp.Keys()
			for _, k := range keys {
				err := field.Add(k, tmp.Get(k))
				if err != nil {
					return field, err
				}
			}
		}
		return field, nil
	}
}
