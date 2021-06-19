package logger

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/cheesycoffee/go-utilities/timeutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger     *zap.SugaredLogger
	production *bool
	once       sync.Once
)

// Init logger
func Init(isProduction bool) {
	once.Do(func() {
		var (
			logg *zap.Logger
			err  error
		)
		prod := isProduction
		production = &prod

		cfg := zap.Config{
			Encoding:         "json",
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey: "caller",
				EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
					_, caller.File, caller.Line, _ = runtime.Caller(7)
					enc.AppendString(caller.FullPath())
				},
			},
			Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
			Development: !*production,
		}

		logg, err = cfg.Build()
		if err != nil {
			panic(err)
		}
		defer logg.Sync()

		// define logger
		logger = logg.Sugar()
	})
}

// Log func
func Log(level zapcore.Level, message string, context string, scope string) {
	entry := logger.With(
		zap.String("context", context),
		zap.String("scope", scope),
	)

	switch level {
	case zapcore.DebugLevel:
		entry.Debug(message)
	case zapcore.InfoLevel:
		entry.Info(message)
	case zapcore.WarnLevel:
		entry.Warn(message)
	case zapcore.ErrorLevel:
		entry.Error(message)
	case zapcore.FatalLevel:
		entry.Fatal(message)
	case zapcore.PanicLevel:
		entry.Panic(message)
	}
}

// LogE error
func LogE(message string) {
	logger.Error(message)
}

// LogEf error with format
func LogEf(format string, i ...interface{}) {
	logger.Errorf(format, i...)
}

// LogI info
func LogI(message string) {
	logger.Info(message)
}

// LogIf info with format
func LogIf(format string, i ...interface{}) {
	logger.Infof(format, i...)
}

// LogWithDefer return defer func for status
func LogWithDefer(str string) (deferFunc func()) {
	fmt.Printf("%s %s ", time.Now().Format(timeutil.TimeFormatLogger), str)
	return func() {
		if r := recover(); r != nil {
			fmt.Printf("\x1b[31;1mERROR: %v\x1b[0m\n", r)
			panic(r)
		}
		fmt.Println("\x1b[32;1mSUCCESS\x1b[0m")
	}
}

// LogYellow log with yellow color
func LogYellow(str string) {
	if !*production {
		fmt.Printf("\x1b[33;5m%s\x1b[0m\n", str)
	}
}

// LogRed log with red color
func LogRed(str string) {
	if !*production {
		fmt.Printf("\x1b[31;5m%s\x1b[0m\n", str)
	}
}

// LogGreen log with green color
func LogGreen(str string) {
	if !*production {
		fmt.Printf("\x1b[32;5m%s\x1b[0m\n", str)
	}
}
