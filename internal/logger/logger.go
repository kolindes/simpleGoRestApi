package logger

import (
	"log"
	"os"
	"time"

	"github.com/kolindes/simpleRestApi/internal/config"
)

// Colors ...
const (
	Reset       = "\033[0m"
	White       = "\033[37m"
	WhiteBold   = "\033[37;1m"
	Cyan        = "\033[36m"
	CyanBold    = "\033[36;1m"
	Green       = "\033[32m"
	GreenBold   = "\033[32;1m"
	Yellow      = "\033[33m"
	YellowBold  = "\033[33;1m"
	Red         = "\033[31m"
	RedBold     = "\033[31;1m"
	Magenta     = "\033[35m"
	MagentaBold = "\033[35;1m"
	Blue        = "\033[34m"
	BlueBold    = "\033[34;1m"
)

const (
	InfoTag  = "I"
	WarnTag  = "W"
	ErrorTag = "E"
	FatalTag = "F"
)

const (
	// 1 -- No logs
	Silent int = iota + 1
	// 2 -- +Fatal
	Fatal
	// 3 -- +Error
	Error
	// 4 -- +Warn
	Warn
	// 5 -- +Info
	Info
)

// Writer ...
type Writer interface {
	Printf(string, ...interface{})
}

type Logger struct {
	Writer
	LogLevel int
	infoStr  string
	warnStr  string
	errStr   string
	fatalStr string
}

// Interface ...
type Interface interface {
	LogMode(int) Interface
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

// New ...
func New(config config.LoggingConfig) Interface {
	writer := log.New(os.Stdout, "", 0)
	var (
		infoStr  = "[%s]\t" + InfoTag + "\t%s"
		warnStr  = "[%s]\t" + WarnTag + "\t%s"
		errorStr = "[%s]\t" + ErrorTag + "\t%s"
		fatalStr = "[%s]\t" + FatalTag + "\t%s"
	)

	if config.Colorful {
		infoStr = White + "[" + Cyan + "%s" + White + "]\t" + Reset + WhiteBold + InfoTag + "\t%s" + Reset
		warnStr = White + "[" + Cyan + "%s" + White + "]\t" + Reset + YellowBold + WarnTag + "\t%s" + Reset
		errorStr = White + "[" + Cyan + "%s" + White + "]\t" + Reset + RedBold + ErrorTag + "\t%s" + Reset
		fatalStr = White + "[" + Cyan + "%s" + White + "]\t" + Reset + MagentaBold + FatalTag + "\t%s" + Reset
	}

	return &Logger{
		Writer:   writer,
		LogLevel: config.LogLevel,
		infoStr:  infoStr,
		warnStr:  warnStr,
		errStr:   errorStr,
		fatalStr: fatalStr,
	}
}

// LogMode ...
func (l *Logger) LogMode(level int) Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info ...
func (l *Logger) Info(msg string) {
	if l.LogLevel >= Info {
		l.Printf(l.infoStr, time.Now().Format(time.RFC3339), msg)
	}
}

// Warn ...
func (l *Logger) Warn(msg string) {
	if l.LogLevel >= Warn {
		l.Printf(l.warnStr, time.Now().Format(time.RFC3339), msg)
	}
}

// Error ...
func (l *Logger) Error(msg string) {
	if l.LogLevel >= Error {
		l.Printf(l.errStr, time.Now().Format(time.RFC3339), msg)
	}
}

// Fatal ...
func (l *Logger) Fatal(msg string) {
	if l.LogLevel >= Fatal {
		l.Printf(l.fatalStr, time.Now().Format(time.RFC3339), msg)
	}
}
