// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"os"
	"sort"
	"testing"
	"time"
)

func TestField(t *testing.T) {
	kvs := xlog.NewKeyValues()
	err := kvs.Add("int", 1, "time", time.Now(), "nil")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("keys", kvs.Keys())
	t.Log("all", kvs.GetAll())
	it := kvs.Iterator()
	for it.HasNext() {
		x, v := it.Next()
		t.Log("iterator", x, v)
	}

	kvs.Add("float", 1.1, "string", "test", "int", -1)
	t.Log("after add twice")
	it = kvs.Iterator()
	for it.HasNext() {
		x, v := it.Next()
		t.Log("iterator", x, v)
	}

	kvs.Remove("float")
	t.Log("after remove float")
	it = kvs.Iterator()
	for it.HasNext() {
		x, v := it.Next()
		t.Log("iterator", x, v)
	}

	clone := kvs.Clone()
	t.Log("after Clone float")
	it = clone.Iterator()
	for it.HasNext() {
		x, v := it.Next()
		t.Log("iterator", x, v)
	}

	for k, v := range clone.GetAll() {
		t.Log("key: ", k, " value: ", v)
	}
}

func TestTextFormatter(t *testing.T) {
	kvs := xlog.NewKeyValues()
	err := kvs.Add("int", 1, "time", time.Now(), "nil", nil, "float", 1.1, "string", "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("empty", func(t *testing.T) {
		formatter := xlog.TextFormatter{}
		err := formatter.Format(os.Stdout, xlog.NewKeyValues())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("none", func(t *testing.T) {
		formatter := xlog.TextFormatter{}
		err := formatter.Format(os.Stdout, kvs)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("sorted", func(t *testing.T) {
		formatter := xlog.TextFormatter{
			SortFunc: sort.Strings,
		}
		err := formatter.Format(os.Stdout, kvs)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("sorted with quota", func(t *testing.T) {
		formatter := xlog.TextFormatter{
			SortFunc:  sort.Strings,
			WithQuote: true,
		}
		err := formatter.Format(os.Stdout, kvs)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestJsonFormatter(t *testing.T) {
	kvs := xlog.NewKeyValues()
	err := kvs.Add("int", 1, "time", time.Now(), "nil", nil, "float", 1.1, "string", "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("none", func(t *testing.T) {
		formatter := xlog.JsonFormatter{}
		err := formatter.Format(os.Stdout, kvs)
		if err != nil {
			t.Fatal(err)
		}
	})
}
