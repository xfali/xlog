// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func TestLoggingf(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l.Logf(xlog.DEBUG, 0, "", "DEBUG test\n")
	l.Logf(xlog.INFO, 0, "", "INFO test\n")
	l.Logf(xlog.WARN, 0, "mytest", "WARN test\n")
	l.Logf(xlog.ERROR, 0, "", "ERROR test")
	l.Logf(xlog.FATAL, 0, "", "FATAL test\n")
}

func TestLogging(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l.Log(xlog.DEBUG, 0, "", "DEBUG", " test")
	l.Log(xlog.INFO, 0, "", "INFO", " test")
	l.Log(xlog.WARN, 0, "", "WARN", " test")
	l.Log(xlog.ERROR, 0, "", "ERROR", " test")
	l.Log(xlog.FATAL, 0, "", "FATAL", " test")
}

func TestLoggingln(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l.Logln(xlog.DEBUG, 0, "", "DEBUG", " test")
	l.Logln(xlog.INFO, 0, "", "INFO", " test")
	l.Logln(xlog.WARN, 0, "", "WARN", " test")
	l.Logln(xlog.ERROR, 0, "", "ERROR", " test")
	l.Logln(xlog.FATAL, 0, "", "FATAL", " test")
}
