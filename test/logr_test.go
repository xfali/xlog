// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"errors"
	"github.com/xfali/xlog"
	"github.com/xfali/xlog/xlogr"
	"testing"
	"time"
)

func TestLogr(t *testing.T) {
	logr := xlogr.NewLogr()
	if !logr.Enabled() {
		t.Fatal("logr is disabled")
	}

	logr = logr.V(xlogr.Level2Int(xlog.DEBUG))
	if logr.Enabled() {
		t.Fatal("logr is enable")
	}

	logr = logr.V(xlogr.Level2Int(xlog.WARN))
	if !logr.Enabled() {
		t.Fatal("logr is disabled")
	}

	logr = logr.WithName("LOGR_TEST")
	logr = logr.WithValues("123", "abc")
	logr.Info("this is a test", "time", time.Now(), "float", 3.14)
	logr.Error(errors.New("test"), "this is a test", "time", time.Now(), "float", 3.14)

	logr = logr.WithName("xxxx")
	logr.Info("this is a test", "time", time.Now(), "float", 3.14)
}

func TestLogr2(t *testing.T) {
	logging := xlog.NewLogging()
	logging.SetFormatter(&xlog.TextFormatter{
		TimeFormat: xlog.TimeFormat,
		WithQuote:  true,
	})
	logr := xlogr.NewLogrWithLogging(logging)
	if !logr.Enabled() {
		t.Fatal("logr is disabled")
	}

	logr = logr.V(xlogr.Level2Int(xlog.DEBUG))
	if logr.Enabled() {
		t.Fatal("logr is enable")
	}

	logr = logr.V(xlogr.Level2Int(xlog.WARN))
	if !logr.Enabled() {
		t.Fatal("logr is disabled")
	}

	logr = logr.WithName("LOGR_TEST")
	logr = logr.WithValues("123", "abc")
	logr.Info("this is a test", "time", time.Now(), "float", 3.14)
	logr.Error(errors.New("test"), "this is a test", "time", time.Now(), "float", 3.14)
}
