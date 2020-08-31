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

type Field interface {
	Add(keyAndValues ...interface{}) error
	GetAll() map[string]interface{}
	Keys() []string
	Get(key string) interface{}
	Iterator() Iterator

	Clone() Field
}

type Formatter interface {
	Format(writer io.Writer, field Field) error
}

type defaultField [2]interface{}

func NewField(keyAndValues ...interface{}) Field {
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

func (f defaultField) Clone() Field {
	ret := NewField()

	for _, k := range f.Keys() {
		ret.Add(k, f.Get(k))
	}

	return ret
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

func MergeFields(fields ...Field) (Field, error) {
	if len(fields) == 0 {
		return nil, errors.New("No field to merge ")
	} else {
		field := fields[0]
		var tmp Field
		for i := 1; i < len(fields); i++ {
			tmp = fields[i]
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

type TextFormatter struct {
	TimeFormat func(t time.Time) string
	WithQuote  bool
	SortFunc   func([]string)
}

func (f *TextFormatter) Format(writer io.Writer, field Field) error {
	keys := field.Keys()
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
		buf.WriteString(f.formatValue(field.Get(k)))
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

func (f *JsonFormatter) Format(writer io.Writer, field Field) error {
	d, err := json.Marshal(field.GetAll())
	if err != nil {
		return err
	}
	_, err = writer.Write(d)
	return err
}
