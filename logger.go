package logger

import (
	"github.com/xu215740578/zap"
	"github.com/xu215740578/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var errorLogger *zap.SugaredLogger

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

//InitLogger set logger info and init logger
func InitLogger(file string, level string, maxsize int, maxage int) *zap.Logger {
	loglevel := getLoggerLevel(level)

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: file,
		MaxSize:  maxsize, // megabytes
		// MaxBackups: 3,
		MaxAge: maxage, // days
	})

	encoderCfg := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), w, zap.NewAtomicLevelAt(loglevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
	return logger
}

//Debugln log Debug level info
func Debugln(args ...interface{}) {
	errorLogger.Debug(args...)
}

//Debugf log Debug level info
func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

//Infoln log Info level info
func Infoln(args ...interface{}) {
	errorLogger.Info(args...)
}

//Infof log Info level info
func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

//Warnln log Warn level info
func Warnln(args ...interface{}) {
	errorLogger.Warn(args...)
}

//Warnf log Warn level info
func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

//Errorln log Error level info
func Errorln(args ...interface{}) {
	errorLogger.Error(args...)
}

//Errorf log Error level info
func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

//DPanicln log Panic level info
func DPanicln(args ...interface{}) {
	errorLogger.DPanic(args...)
}

//DPanicf log Panic level info
func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

//Panicln log Panic level info
func Panicln(args ...interface{}) {
	errorLogger.Panic(args...)
}

//Panicf log Panic level info
func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

//Fatalln log Fatal level info
func Fatalln(args ...interface{}) {
	errorLogger.Fatal(args...)
}

//Fatalf log Fatal level info
func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
