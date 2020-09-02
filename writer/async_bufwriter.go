// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description:

package writer

import (
	"bytes"
	"errors"
	"io"
	"time"
)

const (
	FlushSize = 10240
	FlushTime = 5 * time.Millisecond
)

type AsyncBufferLogWriter struct {
	stopChan  chan bool
	logChan   chan []byte
	logBuffer bytes.Buffer
	FlushSize int64
	w         io.WriteCloser
	block     bool
}

type Config struct {
	// 触发刷新的数据大小阈值
	FlushSize int64

	// 异步缓存的大小，如果超出可能会阻塞或返回错误（由Block控制）
	BufferSize int

	// 触发刷新的时间间隔
	FlushInterval time.Duration

	// 如果为true，则当超出bufSize大小时Write方法阻塞，否则返回error
	Block bool
}

var defaultConfig = Config{
	FlushSize:     FlushSize,
	BufferSize:    BufferSize,
	FlushInterval: FlushTime,
	Block:         true,
}

// 带Buffer的Writer
// Param: w - 实际写入的Writer， c - Writer的配置，如果不传入则使用默认值，否则使用第1个配置。
func NewAsyncBufferWriter(w io.WriteCloser, c ...Config) *AsyncBufferLogWriter {
	conf := defaultConfig
	if len(c) > 0 {
		conf = c[0]
	}
	l := AsyncBufferLogWriter{
		stopChan:  make(chan bool),
		logChan:   make(chan []byte, conf.BufferSize),
		FlushSize: conf.FlushSize,
		w:         w,
		block:     conf.Block,
	}
	l.logBuffer.Grow(conf.BufferSize * 10)

	go func() {
		defer func() {
			if w != nil {
				l.Flush()
				w.Close()
			}
		}()
		ticker := time.NewTicker(conf.FlushInterval)
		defer ticker.Stop()
		for {
			select {
			case <-l.stopChan:
				return
			case d, ok := <-l.logChan:
				if ok {
					l.writeLog(d)
				}
			case <-ticker.C:
				l.Flush()
			}
		}
	}()
	return &l
}

func (w *AsyncBufferLogWriter) Flush() error {
	if w.w != nil {
		_, err := w.w.Write(w.logBuffer.Bytes())
		if err != nil {
			return err
		}
	}
	w.logBuffer.Reset()
	return nil
}

func (w *AsyncBufferLogWriter) writeLog(data []byte) error {
	w.logBuffer.Write(data)

	if int64(w.logBuffer.Len()) < w.FlushSize {
		return nil
	}

	return w.Flush()
}

func (w *AsyncBufferLogWriter) Close() error {
	close(w.stopChan)
	return nil
}

func (w *AsyncBufferLogWriter) Write(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}
	if w.block {
		w.logChan <- data
		return len(data), nil
	} else {
		select {
		case w.logChan <- data:
			return len(data), nil
		default:
			return 0, errors.New("write log failed ")
		}
	}
}
