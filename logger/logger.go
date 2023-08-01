package logger

import (
	"fmt"
	"os"
)

const NAME = "LOG"

const (
	ERROR int = 0
	WARN      = 1
	DEBUG     = 2
	INFO      = 3
)

var LogLevels = map[int]string {
	ERROR: "ERROR",
	WARN: "WARN",
	DEBUG: "DEBUG",
	INFO: "INFO",
}

type logger struct {
	log_level int
}

var l = &logger{}

func Initialise(log_level int) {
	l.log_level = log_level
	LogInf(NAME, "Logger initialised with level " + LogLevels[log_level] + ".")
}

func LogRaw(message string) {
	fmt.Println(message)
}

func LogfRaw(message string, params ...interface{}) {
	fmt.Printf(message, params...)
	fmt.Println()
}

func LogErr(module string, message string) {
	fmt.Fprint(os.Stderr, "[ERR] "+module+": "+message)
	fmt.Println()
}

func LogfErr(module string, message string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERR] "+module+": "+message, params...)
	fmt.Println()
}

func LogWar(module string, message string) {
	if l.log_level >= WARN {
		fmt.Println("[WRN] " + module + ": " + message)
	}
}

func LogfWar(module string, message string, params ...interface{}) {
	if l.log_level >= WARN {
		fmt.Printf("[WRN] "+module+": "+message, params...)
		fmt.Println()
	}
}

func LogDbg(module string, message string) {
	if l.log_level >= DEBUG {
		fmt.Println("[DBG] " + module + ": " + message)
	}
}

func LogfDbg(module string, message string, params ...interface{}) {
	if l.log_level >= DEBUG {
		fmt.Printf("[DBG] "+module+": "+message, params...)
		fmt.Println()
	}
}

func LogfInf(module string, message string, params ...interface{}) {
	if l.log_level >= INFO {
		fmt.Printf("[INF] "+module+": "+message, params...)
		fmt.Println()
	}
}

func LogInf(module string, message string) {
	if l.log_level >= INFO {
		fmt.Println("[INF] " + module + ": " + message)
	}
}
