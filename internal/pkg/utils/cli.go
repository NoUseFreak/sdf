package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func SdfBinaryPath() string {
	if strings.Contains(os.Args[0], "go-build") {
		cwd, _ := os.Getwd()
		return fmt.Sprintf("go run %s", path.Join(cwd, "cmd/sdf/main.go"))
	}

	sdfPath, err := exec.LookPath(os.Args[0])
	if err != nil {
		return os.Args[0]
	}

	return sdfPath
}
