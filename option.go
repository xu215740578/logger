package logger

type Options struct {
	FileName string
	MaxAge   int    // 日志文件保留时长(天)
	MaxSize  int    // 日志文件切割时间(小时)
	Level    string // debug err info
}

type Option func(*Options)

func defaultOption() Options {
	return Options{
		FileName: "./server.log",
		MaxAge:   7,   // 7天
		MaxSize:  100, // 100M
		Level:    "debug",
	}
}
