package infra

import (
	"os"
	"io"
	"runtime"
	"fmt"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger(levelStr string, filePath string) {
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
	Log.SetReportCaller(true)

	Log.SetOutput(io.Discard) // bỏ output mặc định

	Log.AddHook(newHook(os.Stdout, terminalFormatter()))
	Log.AddHook(newHook(openLogFile(filePath), fileFormatter()))
}

func terminalFormatter() logrus.Formatter {
	return &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", f.File, f.Line)
		},
	}
}

func fileFormatter() logrus.Formatter {
	return &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

func openLogFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("❌ Cannot open log file: %v", err))
	}
	return file
}

func newHook(writer io.Writer, formatter logrus.Formatter) logrus.Hook {
	return &writerHook{
		Writer:    writer,
		LogLevels: logrus.AllLevels,
		Formatter: formatter,
	}
}

type writerHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
	Formatter logrus.Formatter
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}