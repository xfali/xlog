// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package xlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type TextFormatter struct {
	TimeFormat func(t time.Time) string
	WithQuote  bool
	SortFunc   func([]string)
}

func (f *TextFormatter) Format(writer io.Writer, field Field) error {
	kvs := ToKeyValues(field)
	keys := kvs.Keys()
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
		buf.WriteString(f.formatValue(kvs.Get(k)))
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
	v := ToKeyValues(field)
	d, err := json.Marshal(v.GetAll())
	if err != nil {
		return err
	}
	_, err = writer.Write(d)
	return err
}
