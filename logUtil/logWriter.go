package logutil

import (
	"fmt"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

type logWriter struct {
	logEndTime  int64
	locker      sync.Mutex
	fs          *os.File
	baseLogPath string
	typeName    string
}

// newLogWriter 创建一个按小时轮转的日志写入器。logPath 为根目录，typeName 作为文件名后缀。
func newLogWriter(logPath string, typeName string) *logWriter {
	return &logWriter{
		baseLogPath: logPath,
		locker:      sync.Mutex{},
		logEndTime:  0,
		typeName:    typeName,
	}
}

// Write 实现 io.Writer，写入日志内容。若当前小时文件不存在或已到轮转时间则自动创建新文件。
func (writer *logWriter) Write(p []byte) (n int, err error) {
	err = writer.newWriter()
	if err != nil {
		return
	}
	n, err = writer.fs.Write(p)
	return
}

// Close 关闭当前日志文件，将 logEndTime 置为 -1 表示已关闭，后续写入将返回错误。
func (writer *logWriter) Close() error {
	writer.locker.Lock()
	defer writer.locker.Unlock()

	if writer.logEndTime == 0 {
		return nil
	}

	err := writer.fs.Close()
	if err != nil {
		return err
	}

	writer.logEndTime = 0
	writer.fs = nil
	return nil
}

// newWriter 检查当前文件是否需要轮转，若超出当前小时的有效期则创建新文件。
// 使用 atomic 快速路径避免每次写入都加锁，加锁后再次校验防止并发重复创建。
func (writer *logWriter) newWriter() error {
	logEndTime := atomic.LoadInt64(&writer.logEndTime)
	if logEndTime < 0 {
		return fmt.Errorf("loger closed")
	}

	t := time.Now()
	tk := t.Unix()
	if tk < logEndTime {
		return nil
	}

	writer.locker.Lock()
	defer writer.locker.Unlock()

	if writer.logEndTime < 0 {
		return fmt.Errorf("loger closed")
	}

	if tk < writer.logEndTime {
		return nil
	}

	if writer.fs != nil {
		if err := writer.fs.Close(); err != nil {
			return err
		}
	}

	y, m, d := t.Date()
	logPath := path.Join(writer.baseLogPath, fmt.Sprintf("%d/%d", y, m))
	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		return err
	}

	logFile := path.Join(logPath, fmt.Sprintf("%d-%d-%d_%d_%s.txt", y, m, d, t.Hour(), writer.typeName))
	var err error
	writer.fs, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, mi, s := t.Clock()
	s = mi*60 + s
	t = t.Add(-time.Second * time.Duration(s)).Add(time.Hour)
	writer.logEndTime = t.Unix()
	return nil
}
