package navigate

import (
	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func Chdir(path string) {
	output.Exec("cd %s", path)
}
