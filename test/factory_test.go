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
	logger := xlog.GetLogger()
	logger.Warnln("test")
	logger = xlog.GetLogger(nil)
	logger.Warnln("test")
	logger = logger.WithName("test2")
	logger.Warnln("test2")
	logger = logger.WithName("test3")
	logger.Warnln("test3")
	logger = logger.WithFields("FieldKey", "FieldValue")
	logger.Warnln("test4")
	logger = logger.WithFields("FieldKey", "FieldValue2")
	logger.Warnln("test5")

	type TestStructInTest struct{}
	logger = xlog.GetLogger(TestStructInTest{})
	logger.Warnln("A")

	logger = xlog.GetLogger(1)
	logger.Warnln("int")
}

func TestFactorySimplifyName(t *testing.T) {
	fac := xlog.NewFactory(xlog.NewLogging())
	fac.SimplifyNameFunc = xlog.SimplifyNameFirstLetter

	xlog.ResetFactory(fac)

	type TestStructInTest struct{}
	logger := xlog.GetLogger(TestStructInTest{})
	logger.Warnln("A")

	logger = xlog.GetLogger(1)
	logger.Warnln("int")
}
