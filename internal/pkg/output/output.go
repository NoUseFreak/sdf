package output

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func Print(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func Println(format string, a ...interface{}) {
	Print(format+"\n", a...)
}

func Error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, aurora.Red(fmt.Sprintf(format, a...)).String())
}

func Errorln(format string, a ...interface{}) {
	Error(format+"\n", a...)
}

func Exec(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", a...)
}

func Debug(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}
