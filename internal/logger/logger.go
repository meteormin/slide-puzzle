package logger

import (
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// New initializes and returns a new sugared logger.
//
// It accepts a variable number of Config options.
// Returns a pointer to a zap.SugaredLogger.
func New(config ...Config) *zap.SugaredLogger {
	cfg := resolveConfig(config...)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = cfg.TimeKey
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		loc, err := time.LoadLocation(cfg.TimeZone)
		if err != nil {
			loc = time.Local
		}

		t = t.In(loc)

		type appendTimeEncoder interface {
			// AppendTimeLayout description of the Go function.
			//
			// This function takes a time.Time and a string as parameters.
			AppendTimeLayout(time.Time, string)
		}

		if enc, ok := enc.(appendTimeEncoder); ok {
			enc.AppendTimeLayout(t, cfg.TimeFormat)
			return
		}

		enc.AppendString(t.Format(cfg.TimeFormat))
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	coreSlice := make([]zapcore.Core, 0)
	consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), cfg.LogLevel)
	coreSlice = append(coreSlice, consoleCore)

	logFilename := path.Join(cfg.FilePath, cfg.Filename)
	if logFilename != "" {
		ll := &lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}

		ws := zapcore.AddSync(ll)
		lLCore := zapcore.NewCore(encoder, ws, cfg.LogLevel)
		coreSlice = append(coreSlice, lLCore)
	}

	core := zapcore.NewTee(coreSlice...)
	zapLogger := zap.New(core, zap.AddCaller())
	logger := zapLogger.Named(cfg.Name).Sugar()

	if cfg.WithOptions != nil {
		logger = logger.WithOptions(cfg.WithOptions...)
	}

	return logger
}
