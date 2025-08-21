package utils

type LogLevel string

const (
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelDebug LogLevel = "DEBUG"
	LevelError LogLevel = "ERROR"
)

func IsValidLevel(level string) bool {
	switch level {
	case string(LevelInfo), string(LevelWarn), string(LevelError):
		return true
	default:
		return false
	}
}
