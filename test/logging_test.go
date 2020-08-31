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
	l.Logf(xlog.DEBUG, 0, nil, "DEBUG test\n")
	l.Logf(xlog.INFO, 0, nil, "INFO test\n")
	l.Logf(xlog.WARN, 0, xlog.NewField("mytest", "mytest"), "WARN test\n")
	l.Logf(xlog.ERROR, 0, nil, "ERROR test")
	l.Logf(xlog.FATAL, 0, nil, "FATAL test\n")
}

func TestLogging(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l.Log(xlog.DEBUG, 0, nil, "DEBUG", " test")
	l.Log(xlog.INFO, 0, nil, "INFO", " test")
	l.Log(xlog.WARN, 0, nil, "WARN", " test")
	l.Log(xlog.ERROR, 0, nil, "ERROR", " test")
	l.Log(xlog.FATAL, 0, nil, "FATAL", " test")
}

func TestLoggingln(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l.Logln(xlog.DEBUG, 0, nil, "DEBUG", " test")
	l.Logln(xlog.INFO, 0, nil, "INFO", " test")
	l.Logln(xlog.WARN, 0, nil, "WARN", " test")
	l.Logln(xlog.ERROR, 0, nil, "ERROR", " test")
	l.Logln(xlog.FATAL, 0, nil, "FATAL", " test")
}

func TestLoggingWithFormat(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetShowFileFlag(xlog.LongFile))
	l := xlog.NewLogging()
	l.SetFormatter(&xlog.TextFormatter{
		TimeFormat: xlog.TimeFormat,
	})
	l.Logln(xlog.ERROR, 0, xlog.NewField("int", 1, "string", "2"), "ERROR", " test")
}
