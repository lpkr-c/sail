package slog

import (
	"fmt"
)

var slogLevel int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

func init() {
	slogLevel = INFO
}

func SetLevel(level int) {
	slogLevel = level
}

func values(level int, fieldValues []interface{}) {
	if level >= slogLevel {
		for i := 0; i < len(fieldValues); i += 2 {
			fmt.Printf("\t%s: %+v\n", fieldValues[i], fieldValues[i+1])
		}
	}
}

func InfoValues(fieldValues ...interface{}) {
	values(INFO, fieldValues)
}

func DebugValues(fieldValues ...interface{}) {
	values(DEBUG, fieldValues)
}

func printf(level int, format string, values ...interface{}) {
	if level >= slogLevel {
		fmt.Printf(format, values...)
	}
}

func ErrorPrintf(format string, values ...interface{}) {
	printf(ERROR, format, values...)
}
func WarnPrintf(format string, values ...interface{}) {
	printf(WARN, format, values...)
}
func InfoPrintf(format string, values ...interface{}) {
	printf(INFO, format, values...)
}

func DebugPrintf(format string, values ...interface{}) {
	printf(DEBUG, format, values...)
}
