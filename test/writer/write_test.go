// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package writer

import (
	"fmt"
	"github.com/xfali/xlog/writer"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAsyncWriter(t *testing.T) {
	w := writer.NewAsyncWriter(os.Stdout, 10, true)
	var count int32 = 0
	wait := sync.WaitGroup{}
	wait.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wait.Done()
			b := make([]byte, 10)
			for i := 0; i < 10; i++ {
				atomic.AddInt32(&count, 1)
				rand.Read(b)
				fmt.Fprintln(w, string(b))
			}
		}()
	}

	wait.Wait()
	t.Log(count)
}

func TestAsyncBufWriter(t *testing.T) {
	w := writer.NewAsyncBufferWriter(os.Stdout, writer.Config{
		FlushSize:     100,
		BufferSize:    10,
		FlushInterval: 1 * time.Millisecond,
		Block:         true,
	})
	var count int32 = 0
	wait := sync.WaitGroup{}
	wait.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wait.Done()
			b := make([]byte, 10)
			for i := 0; i < 10; i++ {
				atomic.AddInt32(&count, 1)
				rand.Read(b)
				fmt.Fprintln(w, string(b))
			}
		}()
	}

	wait.Wait()
	t.Log(count)
}
