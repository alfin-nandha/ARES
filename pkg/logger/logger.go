package logger

import (
	"ares/helper"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	RotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

var Log *Logger

type Logger struct {
	loggerSys *zap.Logger
	mask      []string
}

type Options struct {
	FileLocation string        `json:"fileLocation"`
	FileMaxAge   time.Duration `json:"fileMaxAge"`
	Stdout       bool          `json:"stdout"`
	Mask         []string      `json:"mask"`
}

func New(config Options) *Logger {
	var cores []zapcore.Core

	var writer zapcore.WriteSyncer

	if config.Stdout {
		writer = zapcore.AddSync(os.Stdout)
	} else {
		rotate, err := RotateLogs.New(
			config.FileLocation+".%Y-%m-%d",
			RotateLogs.WithMaxAge(config.FileMaxAge*24*time.Hour),
			RotateLogs.WithRotationTime(time.Hour),
		)
		if err != nil {
			panic(err)
		}
		writer = zapcore.AddSync(rotate)
	}

	core := zapcore.NewCore(getEncoder(), writer, zapcore.InfoLevel)
	cores = append(cores, core)

	combinedCore := zapcore.NewTee(cores...)

	loggerSys := zap.New(combinedCore,
		zap.AddCallerSkip(3),
		zap.AddCaller(),
	)

	// return &Logger{
	// 	loggerSys: loggerSys,
	// 	mask:      config.Mask,
	// }
	Log = &Logger{
		loggerSys: loggerSys,
		mask:      config.Mask,
	}
	return Log
}

func getEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "message",
		EncodeDuration: MillisDurationEncoder,
		EncodeTime:     TDRLogTimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}
	return zapcore.NewConsoleEncoder(config)
}

func TDRLogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	enc.AppendString(t.In(location).Format("2006-01-02 15:04:05.999"))
}

func MillisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Milliseconds())
}

func (l *Logger) Error(message string, fields ...zap.Field) {
	l.loggerSys.Error(message, fields...)
}

func (l *Logger) Info(message string, fields ...zap.Field) {
	l.loggerSys.Info(message, fields...)
}

// func (l *Logger) InfoLite(message string, fields ...zap.Field) {
// 	l.loggerTdr.Info(message, fields...)
// }

func (l *Logger) MaskData(request interface{}) interface{} {
	data := map[string]interface{}{}
	helper.ObjectToObject(request, &data)

	for _, key := range l.mask {
		if data[key] != nil {
			data[key] = "*****"
		}
	}
	return data
}
