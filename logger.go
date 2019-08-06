package logger

import (
	"github.com/xu215740578/zap"
	"github.com/xu215740578/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var errorLogger *zap.SugaredLogger
var logger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

//InitLogger 设置日志基本配置
func InitLogger(file string, level string, maxsize int, maxage int) {
	loglevel := getLoggerLevel(level)

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: file,
		MaxSize:  maxsize, // megabytes
		// MaxBackups: 3,
		MaxAge: maxage, // days
	})

	encoderCfg := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Name",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "Stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), w, zap.NewAtomicLevelAt(loglevel))
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

//Debugln 打印"Debug"级别日志信息
func Debugln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Debug(args...)
	}
}

//Debugf 打印"Debug"级别日志信息
func Debugf(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Debugf(template, args...)
	}
}

//Infoln 打印"Info"级别日志信息
func Infoln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Info(args...)
	}
}

//Infof 打印"Info"级别日志信息
func Infof(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Infof(template, args...)
	}
}

//Warnln 打印"Warn"级别日志信息
func Warnln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Warn(args...)
	}
}

//Warnf 打印"Warn"级别日志信息
func Warnf(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Warnf(template, args...)
	}
}

//Errorln 打印"Error"级别日志信息
func Errorln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Error(args...)
	}
}

//Errorf 打印"Error"级别日志信息
func Errorf(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Errorf(template, args...)
	}
}

//DPanicln 打印"Panic"级别日志信息
func DPanicln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.DPanic(args...)
	}
}

//DPanicf 打印"Panic"级别日志信息
func DPanicf(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.DPanicf(template, args...)
	}
}

//Panicln 打印"Panic"级别日志信息
func Panicln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Panic(args...)
	}
}

//Panicf 打印"Panic"级别日志信息
func Panicf(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Panicf(template, args...)
	}
}

//Fatalln 打印"Fatal"级别日志信息
func Fatalln(args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Fatal(args...)
	}
}

//Fatalf 打印"Fatal"级别日志信息
func Fatalf(template string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Fatalf(template, args...)
	}
}
