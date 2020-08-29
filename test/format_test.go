// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"sort"
	"testing"
	"time"
)

func TestField(t *testing.T) {
	field := xlog.NewField()
	err := field.Add("int", 1, "time", time.Now(), "nil")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("keys", field.Keys())
	t.Log("all", field.GetAll())
	it := field.Iterator()
	for  it.HasNext() {
		x, v := it.Next()
		t.Log("iterator", x, v)
	}

	field.Add("float", 1.1, "string", "test")
	t.Log("after add twice")
	it = field.Iterator()
	for  it.HasNext() {
		x, v := it.Next()
		t.Log("iterator", x, v)
	}
}

func TestTextFormatter(t *testing.T) {
	field := xlog.NewField()
	err := field.Add("int", 1, "time", time.Now(), "nil", nil, "float", 1.1, "string", "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("none", func(t *testing.T) {
		formatter := xlog.TextFormatter{}
		d, err := formatter.Format(field)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(d))
	})

	t.Run("sorted", func(t *testing.T) {
		formatter := xlog.TextFormatter{
			SortFunc: sort.Strings,
		}
		d, err := formatter.Format(field)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(d))
	})
}

func TestJsonFormatter(t *testing.T) {
	field := xlog.NewField()
	err := field.Add("int", 1, "time", time.Now(), "nil", nil, "float", 1.1, "string", "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("none", func(t *testing.T) {
		formatter := xlog.JsonFormatter{}
		d, err := formatter.Format(field)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(d))
	})
}
