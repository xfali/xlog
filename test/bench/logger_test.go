// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package bench

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/xfali/xlog"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	data := make([]string, 3)
	d := [64]byte{}
	for i := 0; i < 3; i++ {
		rand.Read(d[:])
		data[i] = base64.StdEncoding.EncodeToString(d[:])
	}
	var count int32 = 0
	wait := sync.WaitGroup{}
	wait.Add(20)
	logger := xlog.GetLogger()
	for i := 0; i < 20; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < b.N; i++ {
				atomic.AddInt32(&count, 1)
				logger.Infoln(count, "========", string(data[0]), string(data[1]), string(data[2]))
			}
		}()
	}

	wait.Wait()
	b.Log(count)
}

func BenchmarkMutableLogger(b *testing.B) {
	data := make([]string, 3)
	d := [64]byte{}
	for i := 0; i < 3; i++ {
		rand.Read(d[:])
		data[i] = base64.StdEncoding.EncodeToString(d[:])
	}
	var count int32 = 0
	wait := sync.WaitGroup{}
	wait.Add(20)
	xlog.ResetFactory(xlog.NewMutableFactory(xlog.DefaultLogging()))
	logger := xlog.GetLogger()
	for i := 0; i < 20; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < b.N; i++ {
				atomic.AddInt32(&count, 1)
				logger.Infoln(count, "========", string(data[0]), string(data[1]), string(data[2]))
			}
		}()
	}

	wait.Wait()
	b.Log(count)
}
