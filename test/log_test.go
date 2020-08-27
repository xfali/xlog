// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/xlog"
	"testing"
)

func TestLog(t *testing.T) {
	xlog.Infof("%d %d %d %d\n", 1, 2, 3, 4)
	xlog.Infoln(1, 2, 3, 4)
	xlog.Info(1, 2, 3, 4)
}
