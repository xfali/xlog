// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestLoggingf(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetCallerFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetCallerFlag(xlog.CallerLongFile), xlog.SetExitFunc(func(i int) {
		t.Log("exit: ", i)
	}))
	l.Logf(xlog.DEBUG, 0, nil, "DEBUG test\n")
	l.Logf(xlog.INFO, 0, nil, "INFO test\n")
	l.Logf(xlog.WARN, 0, xlog.NewKeyValues("mytest", "mytest"), "WARN test\n")
	l.Logf(xlog.ERROR, 0, nil, "ERROR test")
	l.Logf(xlog.PANIC, 0, nil, "PANIC test")
	l.Logf(xlog.FATAL, 0, nil, "FATAL test\n")
}

func TestLogging(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetCallerFlag(xlog.LongFile))
	l := xlog.NewLogging(xlog.SetCallerFlag(xlog.CallerLongFile|xlog.CallerShortFunc), xlog.SetExitFunc(func(i int) {
		t.Log("exit: ", i)
	}))
	b(l)
	l.Log(xlog.DEBUG, 0, nil, "DEBUG", " test")
	l.Log(xlog.INFO, 0, nil, "INFO", " test")
	l.Log(xlog.WARN, 0, nil, "WARN", " test")
	l.Log(xlog.ERROR, 0, nil, "ERROR", " test")
	l.Log(xlog.FATAL, 0, nil, "FATAL", " test")
}

func TestLoggingln(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetCallerFlag(xlog.LongFile))
	t.Run("long func", func(t *testing.T) {
		l := xlog.NewLogging(xlog.SetCallerFlag(xlog.CallerLongFile|xlog.CallerLongFunc), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		}))
		l.Logln(xlog.DEBUG, 0, nil, "DEBUG", " test")
		l.Logln(xlog.INFO, 0, nil, "INFO", " test")
		l.Logln(xlog.WARN, 0, nil, "WARN", " test")
		l.Logln(xlog.ERROR, 0, nil, "ERROR", " test")
		l.Logln(xlog.FATAL, 0, nil, "FATAL", " test")
	})

	t.Run("simple func", func(t *testing.T) {
		l := xlog.NewLogging(xlog.SetCallerFlag(xlog.CallerLongFile|xlog.CallerSimpleFunc), xlog.SetExitFunc(func(i int) {
			t.Log("exit: ", i)
		}))
		l.Logln(xlog.DEBUG, 0, nil, "DEBUG", " test")
		l.Logln(xlog.INFO, 0, nil, "INFO", " test")
		l.Logln(xlog.WARN, 0, nil, "WARN", " test")
		l.Logln(xlog.ERROR, 0, nil, "ERROR", " test")
		l.Logln(xlog.FATAL, 0, nil, "FATAL", " test")
	})
}

func TestLoggingWithFormat(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetCallerFlag(xlog.LongFile))
	l := xlog.NewLogging()
	l.SetFormatter(&xlog.TextFormatter{
		TimeFormat: xlog.TimeFormat,
	})
	l.Logln(xlog.ERROR, 0, xlog.NewKeyValues("int", 1, "string", "2"), "ERROR", " test")
}

func TestLoggingWithCloneAndFormat(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetCallerFlag(xlog.LongFile))
	l := xlog.NewLogging()
	l.SetFormatter(&xlog.TextFormatter{
		TimeFormat: xlog.TimeFormat,
	})
	l.Logln(xlog.ERROR, 0, xlog.NewKeyValues("int", 1, "string", "2"), "ERROR", " test")
	l = l.Clone()
	l.Logln(xlog.ERROR, 0, xlog.NewKeyValues("int", 1, "string", "2"), "Clone ERROR", " test")
}

func TestLoggingFatal(t *testing.T) {
	go func() {
		time.Sleep(3 * time.Second)
		xlog.Fatalln("test")
	}()
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(100 * time.Millisecond)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func TestLoggingHook(t *testing.T) {
	//l := xlog.NewLogging(xlog.SetCallerFlag(xlog.LongFile))
	l := xlog.NewLogging()
	l = xlog.NewHookLevelLogging(l, func(level xlog.Level) xlog.Level {
		return level + 1
	})
	// not print
	l.Logln(xlog.DEBUG, 0, nil, "DEBUG", " test")
	// not print
	l.Logln(xlog.INFO, 0, nil, "INFO", " test")
	// lv hook -> INFO
	l.Logln(xlog.WARN, 0, nil, "WARN", " test")
	// lv hook -> WARN
	l.Logln(xlog.ERROR, 0, nil, "ERROR", " test")
	// lv hook -> ERROR
	l.Logln(xlog.PANIC, 0, nil, "PANIC", " test")
	// lv hook -> PANIC
	l.Logln(xlog.FATAL, 0, nil, "FATAL", " test")
}

func b(logging xlog.Logging) {
	logging.Logln(xlog.INFO, 0, nil, "test")
}
