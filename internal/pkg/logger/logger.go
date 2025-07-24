package logger

import (
	"fmt"
	"nso-server/internal/pkg/config"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger // optional: global logger

func InitZapLogger(cfg *config.Config) (*zap.SugaredLogger, error) {
	level := zapcore.InfoLevel
	switch cfg.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // ✅ in màu
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	logFile := openLogFile("logs/nso.log")

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), level),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar(), nil
}

func openLogFile(path string) *os.File {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Cannot create log directory: %v\n", err)
		return os.Stdout
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Cannot open log file: %v\n", err)
		return os.Stdout
	}
	return file
}

// ✅ Export các phương thức tiện lợi
func Infof(format string, args ...any) {
	Log.Infof(format, args...)
}

func Info(args ...any) {
	Log.Info(args...)
}

func Warnf(format string, args ...any) {
	Log.Warnf(format, args...)
}

func Warn(args ...any) {
	Log.Warn(args...)
}

func Errorf(format string, args ...any) {
	Log.Errorf(format, args...)
}

func Error(args ...any) {
	Log.Error(args...)
}

func Debugf(format string, args ...any) {
	Log.Debugf(format, args...)
}

func Debug(args ...any) {
	Log.Debug(args...)
}

func Fatalf(format string, args ...any) {
	Log.Fatalf(format, args...)
}

func Fatal(args ...any) {
	Log.Fatal(args...)
}

func Printf(format string, args ...any) {
	Log.Infof(format, args...)
}

func Print(args ...any) {
	Log.Info(args...)
}

func WithError(err error) *zap.SugaredLogger {
	return Log.With("error", err)
}

func WithField(key string, value any) *zap.SugaredLogger {
	return Log.With(key, value)
}
