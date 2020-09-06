// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func TestDefaultInfo(t *testing.T) {
	l := xlog.GetLogger()
	l.Info("1")
	l.Infoln("2")
	l.Infof("%d\n", 3)
}

func TestDefaultDepth(t *testing.T) {
	l := xlog.GetLogger()
	l.Infoln("1")
	l = l.WithDepth(1)
	l.Infoln("2")
	l = l.WithDepth(-2)
	l.Infoln("3")
}
