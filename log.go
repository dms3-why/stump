package stump

import (
	"fmt"
	"os"
	"strings"
)

var Verbose bool

var ErrorPrefix = "ERROR: "

func Error(args ...interface{}) {
	log(ErrorPrefix, args)
}

func Fatal(args ...interface{}) {
	Error(args...)
	os.Exit(1)
}

func Log(args ...interface{}) {
	log("", args)
}

func VLog(args ...interface{}) {
	if Verbose {
		log("", args)
	}
}

func log(prefix string, args []interface{}) {
	prefix = strings.TrimRight(prefix, "\t \n")
	writelog := func(format string, args ...interface{}) {
		n := strings.Count(format, "%")
		if n < len(args) {
			format += strings.Repeat(" %s", len(args)-n)
		}
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		fmt.Printf(format, args...)
	}

	if len(args) == 0 {
		writelog(prefix)
		return
	}

	switch s := args[0].(type) {
	case string:
		writelog(prefix+s, args[1:]...)
	case fmt.Stringer:
		writelog(prefix+s.String(), args[1:]...)
	default:
		format := strings.Repeat("%s ", len(args))
		writelog(prefix+format, args...)
	}
}
