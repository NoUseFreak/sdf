package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func SdfBinaryPath() string {
	sdfPath := os.Args[0]
	if strings.Contains(sdfPath, "go-build") {
		cwd, _ := os.Getwd()
		sdfPath = fmt.Sprintf("go run %s", path.Join(cwd, "cmd/sdf/main.go"))
	}

	return sdfPath
}
