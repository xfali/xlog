// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func TestLogf(t *testing.T) {
	l := xlog.NewLogging()
	l.Logf(xlog.DEBUG, 0, "DEBUG test\n")
	l.Logf(xlog.INFO, 0, "INFO test\n")
	l.Logf(xlog.WARN, 0, "WARN test\n")
	l.Logf(xlog.ERROR, 0, "ERROR test\n")
	l.Logf(xlog.FATAL, 0, "FATAL test\n")
}
