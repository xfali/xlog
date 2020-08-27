// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func TestFactory(t *testing.T) {
	logger := xlog.GetLogger("test")
	logger.Infof("this is a %s test\n", "infof")
	logger.Infoln("this is a infoln test")
	logger.Info("this is a info test\n")
}

func TestFactoryTag(t *testing.T) {
	logger := xlog.GetLogger("test")
	logger.Warnln("test")
	logger.WithName("test2")
	logger.Warnln("test2")

	type A struct {}
	logger = xlog.GetLogger(A{})
	logger.Warnln("A")

	logger = xlog.GetLogger(1)
	logger.Warnln("int")
}
