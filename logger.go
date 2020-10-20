package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger *zap.Logger
	out    io.Writer
	sugar  *zap.SugaredLogger
}

func NewLogger(opts ...Option) *Logger {
	options := defaultOption()
	for _, o := range opts {
		o(&options)
	}

	logger := new(Logger)

	logger.out = &lumberjack.Logger{
		Filename: options.FileName,
		MaxSize:  options.MaxSize, // megabytes
		MaxAge:   options.MaxAge,  // days
	}

	ws := zapcore.AddSync(logger.out)

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

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), ws, zap.NewAtomicLevelAt(getLoggerLevel(options.Level)))
	logger.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger.sugar = logger.logger.Sugar()
	return logger
}

var std = NewLogger()

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

func InitLogger(opts ...Option) {
	std = NewLogger(opts...)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() {
	std.logger.Sync()
}

func Writer() io.Writer {
	return std.out
}

//Debug 打印"Debug"级别日志信息
func Debug(v ...interface{}) {
	std.sugar.Debug(v...)
}

//Debugf 打印"Debug"级别日志信息
func Debugf(format string, v ...interface{}) {
	std.sugar.Debugf(format, v...)
}

//Info 打印"Info"级别日志信息
func Info(v ...interface{}) {
	std.sugar.Info(v...)
}

//Infof 打印"Info"级别日志信息
func Infof(format string, v ...interface{}) {
	std.sugar.Infof(format, v...)
}

//Warn 打印"Warn"级别日志信息
func Warn(v ...interface{}) {
	std.sugar.Warn(v...)
}

//Warnf 打印"Warn"级别日志信息
func Warnf(format string, v ...interface{}) {
	std.sugar.Warnf(format, v...)
}

//Error 打印"Error"级别日志信息
func Error(v ...interface{}) {
	std.sugar.Error(v...)
}

//Errorf 打印"Error"级别日志信息
func Errorf(format string, v ...interface{}) {
	std.sugar.Errorf(format, v...)
}

// Print 打印info级别日志
func Print(v ...interface{}) {
	std.sugar.Info(v...)
}

// Print 打印info级别日志
func Printf(format string, v ...interface{}) {
	std.sugar.Infof(format, v...)
}

// Print 打印info级别日志
func Println(v ...interface{}) {
	std.sugar.Info(v...)
}

// Panic 打印"Panic"级别日志信息
func Panic(v ...interface{}) {
	std.sugar.Panic(v...)
}

//Panicln 打印"Panic"级别日志信息
func Panicln(v ...interface{}) {
	std.sugar.Panic(v...)
}

//Panicf 打印"Panic"级别日志信息
func Panicf(format string, v ...interface{}) {
	std.sugar.Panicf(format, v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.sugar.Fatal(v...)
}

//Fatalln 打印"Fatal"级别日志信息
func Fatalln(v ...interface{}) {
	std.sugar.Fatal(v...)
}

//Fatalf 打印"Fatal"级别日志信息
func Fatalf(format string, v ...interface{}) {
	std.sugar.Fatalf(format, v...)
}
