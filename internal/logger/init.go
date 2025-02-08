package logger

var Log *LoggerWrapper // Синглтон

func init() {
	Log, _ = NewLoggerWrapper()
}
