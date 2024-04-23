package logging

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// writerHook implements logrus.Hook.
// It is a hook that writes logs of specified LogLevels to specified Writer
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook.
// It will format logbook entry to string and write it to appropriate writer
func (w *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range w.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return err
}

// Levels define on which logbook levels this hook will be triggered
func (w *writerHook) Levels() []logrus.Level {
	return w.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

// Init initializes the logger
func Init(_ context.Context) error {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s:%d", filename, frame.Line), fmt.Sprintf("%s()", frame.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)

	if err != nil || os.IsExist(err) {
		return errors.New(fmt.Sprintf("Failed to create logs directory: %s", err))
	} else {
		allLogsFile, errOpenFile := os.OpenFile("logs/all.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if errOpenFile != nil {
			return errors.New(fmt.Sprintf("Failed to open logs file: %s", errOpenFile))
		}

		levels := []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel}

		l.SetOutput(io.Discard)
		l.AddHook(
			&writerHook{
				Writer:    []io.Writer{allLogsFile},
				LogLevels: levels,
			},
		)
	}

	e = logrus.NewEntry(l)

	return nil
}

func GetLogger() *Logger {
	return &Logger{
		Entry: e,
	}
}

func (l *Logger) GetLoggerWithField(key string, value interface{}) *Logger {
	return &Logger{
		l.WithField(key, value),
	}
}
