// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func TestLoggerln(t *testing.T) {
	// reset default init at first
	// no fatal trace, do not exit
	xlog.Init(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
		t.Log("exit: ", i)
	})))
	log := xlog.New()
	defer func() {
		recover()
		t.Log("panic !")
		log.Fatalln("this ia a Fatalln test")
	}()
	log.Debugln("this is a Debugln test")
	log.Infoln("this is a Infoln test")
	log.Warnln("this ia a Warnln test")
	log.Errorln("this ia a Errorln test")
	log.Panicln("this ia a Panicln test")
}

func TestLoggerf(t *testing.T) {
	// reset default init at first
	// no fatal trace, do not exit
	xlog.Init(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
		t.Log("exit: ", i)
	})))
	log := xlog.New()
	defer func() {
		recover()
		t.Log("panic !")
		log.Fatalf("this ia a Fatalf test")
	}()
	log.Debugf("this is a Debugf test")
	log.Infof("this is a Infof test")
	log.Warnf("this ia a Warnf test")
	log.Errorf("this ia a Errorf test")
	log.Panicf("this ia a Panicf test")
}

func TestLogger(t *testing.T) {
	// reset default init at first
	// no fatal trace, do not exit
	xlog.Init(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
		t.Log("exit: ", i)
	})))
	log := xlog.New()
	defer func() {
		recover()
		t.Log("panic !")
		log.Fatal("this ia a Fatal test")
	}()
	log.Debug("this is a Debug test")
	log.Info("this is a Info test")
	log.Warn("this ia a Warn test")
	log.Error("this ia a Error test")
	log.Panic("this ia a Panic test")
}
