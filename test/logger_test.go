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
	t.Run("reset logging", func(t *testing.T) {
		xlog.ResetFactory(xlog.NewFactory(
			xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
				t.Log("exit: ", i)
			}))))
		log := xlog.GetLogger()

		log.Infoln("test")

		logging := xlog.NewLogging()
		logging.SetSeverityLevel(xlog.WARN)
		xlog.ResetLogging(logging)
		log.Infoln("test2")

		log2 := xlog.GetLogger()
		if log2.InfoEnabled() {
			t.Fatal("cannot info")
		}
		log2.Info("cannot be here")
		if !log2.WarnEnabled() {
			t.Fatal("can be warn")
		}
		log2.Warnln("must be here")
	})

	t.Run("default", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		})))
		log := xlog.GetLogger()
		defer func() {
			v := recover()
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("recover panic !", kvs.GetAll())
			}
			log.Fatalln("this ia a Fatalln test")
		}()
		log.Debugln("this is a Debugln test")
		log.Infoln("this is a Infoln test")
		log.Warnln("this ia a Warnln test")
		log.Errorln("this ia a Errorln test")
		log.Panicln("this ia a Panicln test")
	})

	t.Run("exit and panic", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		}), xlog.SetPanicFunc(func(v interface{}) {
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("panic !", kvs.GetAll())
			}
		})))
		log := xlog.GetLogger()
		log.Debugln("this is a Debugln test")
		log.Infoln("this is a Infoln test")
		log.Warnln("this ia a Warnln test")
		log.Errorln("this ia a Errorln test")
		log.Panicln("this ia a Panicln test")
		log.Fatalln("this ia a Fatalln test")
	})
}

func TestLoggerf(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		})))
		log := xlog.GetLogger()
		xlog.SetSeverityLevel(xlog.DEBUG)
		defer func() {
			v := recover()
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("panic !", kvs.GetAll())
			}
			log.Fatalf("recover this ia a Fatalf test")
		}()
		log.Debugf("this is a Debugf test")
		log.Infof("this is a Infof test")
		log.Warnf("this ia a Warnf test")
		log.Errorf("this ia a Errorf test")
		log.Panicf("this ia a Panicf test")
		log.Fatalf("this ia a Fatalf test")
	})

	t.Run("exit and panic", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		}), xlog.SetPanicFunc(func(v interface{}) {
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("panic !", kvs.GetAll())
			}
		})))
		log := xlog.GetLogger()
		log.Debugf("this is a Debugf test")
		log.Infof("this is a Infof test")
		log.Warnf("this ia a Warnf test")
		log.Errorf("this ia a Errorf test")
		log.Panicf("this ia a Panicf test")
		log.Fatalf("this ia a Fatalln test")
	})
}

func TestLogger(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		})))
		log := xlog.GetLogger()
		defer func() {
			v := recover()
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("recover panic !", kvs.GetAll())
			}
			log.Fatal("this ia a Fatal test")
		}()
		log.Debug("this is a Debug test")
		log.Info("this is a Info test")
		log.Warn("this ia a Warn test")
		log.Error("this ia a Error test")
		log.Panic("this ia a Panic test")
	})

	t.Run("exit and exit", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		}), xlog.SetPanicFunc(func(v interface{}) {
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("panic !", kvs.GetAll())
			}
		})))
		log := xlog.GetLogger()
		log.Debug("this is a Debug test")
		log.Info("this is a Info test")
		log.Warn("this ia a Warn test")
		log.Error("this ia a Error test")
		log.Panic("this ia a Panic test")
		log.Fatal("this ia a Fatal test")
	})
}

func TestMutableLoggerln(t *testing.T) {
	t.Run("reset logging", func(t *testing.T) {
		xlog.ResetFactory(xlog.NewMutableFactory(
			xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
				t.Log("exit: ", i)
			}))))
		log := xlog.GetLogger()

		log.Infoln("test")

		logging := xlog.NewLogging()
		logging.SetSeverityLevel(xlog.WARN)
		xlog.ResetLogging(logging)
		log.Infoln("cannot be here")

		if log.InfoEnabled() {
			t.Fatal("cannot info")
		}
		log.Info("cannot be here")
		if !log.WarnEnabled() {
			t.Fatal("can be warn")
		}
		log.Warnln("must be here")
	})

	t.Run("default", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetFactory(xlog.NewMutableFactory(
			xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
				t.Log("exit: ", i)
			}))))
		log := xlog.GetLogger()
		defer func() {
			v := recover()
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("recover panic !", kvs.GetAll())
			}
			log.Fatalln("this ia a Fatalln test")
		}()
		log.Debugln("this is a Debugln test")
		log.Infoln("this is a Infoln test")
		log.Warnln("this ia a Warnln test")
		log.Errorln("this ia a Errorln test")
		log.Panicln("this ia a Panicln test")
	})

	t.Run("exit and panic", func(t *testing.T) {
		// reset default init at first
		// no fatal trace, do not exit
		xlog.ResetLogging(xlog.NewLogging(xlog.SetFatalNoTrace(true), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		}), xlog.SetPanicFunc(func(v interface{}) {
			if kvs, ok := v.(xlog.KeyValues); ok {
				t.Log("panic !", kvs.GetAll())
			}
		})))
		log := xlog.GetLogger()
		log.Debugln("this is a Debugln test")
		log.Infoln("this is a Infoln test")
		log.Warnln("this ia a Warnln test")
		log.Errorln("this ia a Errorln test")
		log.Panicln("this ia a Panicln test")
		log.Fatalln("this ia a Fatalln test")
	})
}
