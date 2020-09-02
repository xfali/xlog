// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package writer

import (
	"encoding/base64"
	"github.com/xfali/xlog/writer"
	"math/rand"
	"os"
	"strconv"
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
				_, err := w.Write([]byte(strconv.Itoa(int(count))+base64.StdEncoding.EncodeToString(b) + "\n"))
				if err != nil {
					t.Fatal(err)
				}
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
				_, err := w.Write([]byte(strconv.Itoa(int(count))+base64.StdEncoding.EncodeToString(b) + "\n"))
				if err != nil {
					t.Fatal(err)
				}
			}
		}()
	}

	wait.Wait()
	w.Close()
	t.Log(count)
}

func TestRotateFile(t *testing.T) {
	f := writer.RotateFile{
		MaxFileSize: 10,
	}
	err := f.Open("./test/test.log")
	if err != nil {
		t.Fatal(err)
	}
}
