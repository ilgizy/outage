package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer   []io.Writer
	LogLevel []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevel
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)

	if os.IsNotExist(err) {
		err = os.Mkdir("logs", 0664)
		if err != nil {
			log.Fatal(err)
		}
		allFile, err = os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			log.Fatal(err)
		}
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:   []io.Writer{allFile, os.Stdout},
		LogLevel: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
