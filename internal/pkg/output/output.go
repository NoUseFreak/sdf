package output

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func Print(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func Error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, aurora.Red(fmt.Sprintf(format, a...)).String())
}

func Exec(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", a...)
}
