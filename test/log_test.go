// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func a() {
	xlog.Infoln("test")
}

func TestLog(t *testing.T) {
	a()
	xlog.Infof("%d %d %d %d\n", 1, 2, 3, 4)
	xlog.Infoln(1, 2, 3, 4)
	xlog.Info(1, 2, 3, 4)
}

func TestLogPanic(t *testing.T) {
	defer func() {
		v := recover()
		if kvs, ok := v.(xlog.KeyValues); ok {
			t.Log("recover panic !", kvs.GetAll())
		} else {
			t.Fatal("panic type is not string")
		}
	}()
	xlog.Panic("this ", "is ", "a ", "test")
}
