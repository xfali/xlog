// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package writer

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type RotateFile struct {
	MaxFileSize int64
	DayRotate   bool
	RotateFunc  func(dir string, name string, files ...string) error

	timer      *time.Timer
	fileName   string
	dir        string
	file       *os.File
	curSize    int64
	part       int
	curTimeStr string
}

func (f *RotateFile) Open(filePath string) error {
	if f.MaxFileSize == 0 {
		// no limit
		f.MaxFileSize = math.MaxInt64
	}
	dir := filepath.Dir(filePath)
	_, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	f.dir = dir
	f.fileName = filepath.Base(filePath)

	f.file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	if f.DayRotate {
		f.curTimeStr = time.Now().Format("2006-01-02")
		t := f.nextDay()
		duration := t.Sub(time.Now())
		if duration < 0 {
			duration = 1
		}
		f.timer = time.NewTimer(duration)
	}
	return f.calcPart()
}

func (f *RotateFile) Write(data []byte) (int, error) {
	if f.timer != nil {
		select {
		case <-f.timer.C:
			err := f.rotateDay()
			if err != nil {
				return 0, err
			}
		default:
		}
	}
	f.curSize += int64(len(data))
	if f.curSize >= f.MaxFileSize {
		err := f.rotatePart()
		if err != nil {
			return 0, err
		}
	}
	return f.file.Write(data)
}

func (f *RotateFile) rotateDay() error {
	err := f.changeFile(fmt.Sprintf("%s-%s", f.curTimeStr, f.fileName))
	if err != nil {
		return err
	}
	if f.RotateFunc != nil {
		oldTimeStr := f.curTimeStr
		files, err := ioutil.ReadDir(f.dir)
		if err != nil {
			return err
		}
		var partsFiles []string
		for _, v := range files {
			i := strings.Index(v.Name(), oldTimeStr)
			if i != -1 {
				partsFiles = append(partsFiles, filepath.Join(f.dir, v.Name()))
			}
		}
		if len(partsFiles) > 0 {
			err = f.RotateFunc(f.dir, oldTimeStr+"-"+f.fileName, partsFiles...)
			if err != nil {
				return err
			}
		}
	}

	f.curTimeStr = time.Now().Format("2006-01-02")
	return nil
}

func (f *RotateFile) calcPart() error {
	files, err := ioutil.ReadDir(f.dir)
	if err != nil {
		return err
	}
	part := 0
	prefix := ""
	if f.curTimeStr == "" {
		prefix = "part"
	} else {
		prefix = f.curTimeStr + "-part"
	}
	for _, v := range files {
		i := strings.Index(v.Name(), prefix)
		if i != -1 {
			part++
		}
	}
	f.part = part
	return nil
}

func (f *RotateFile) rotatePart() error {
	if f.curTimeStr == "" {
		return f.changeFile(fmt.Sprintf("part%d-%s", f.part, f.fileName))
	} else {
		return f.changeFile(fmt.Sprintf("%s-part%d-%s", f.curTimeStr, f.part, f.fileName))
	}
}

func (f *RotateFile) changeFile(filename string) error {
	err := f.file.Close()
	if err != nil {
		return err
	}
	err = os.Rename(filepath.Join(f.dir, f.fileName), filepath.Join(f.dir, filename))
	if err != nil {
		return err
	}
	f.file, err = os.OpenFile(filepath.Join(f.dir, f.fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *RotateFile) nextDay() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	t.AddDate(0, 0, 1)
	return t
}

func (f *RotateFile) Close() error {
	if f.timer != nil {
		f.timer.Stop()
	}
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

func ZipLogs(dir string, name string, files ...string) error {
	if len(files) == 0 {
		return nil
	}
	zipFile := filepath.Join(dir, name+".zip")
	f, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer f.Close()
	w := zip.NewWriter(f)
	for _, v := range files {
		err := compress(v, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file string, w *zip.Writer) error {
	of, err := os.Open(file)
	if err != nil {
		return err
	}
	defer of.Close()

	info, err := of.Stat()
	if err != nil {
		return err
	}
	if !info.IsDir() {
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		wh, err :=w.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(wh, of)
		if err != nil {
			return err
		}
	}
	return nil
}
