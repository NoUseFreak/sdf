package env

import (
	"strings"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func Use(name string) error {
	m, err := readModelFile(modelPath(name))
	if err != nil {
		return err
	}

	output.Exec("set -a")
	var keys []string
	for _, l := range strings.Split(m.Data, "\n") {
		parts := strings.Split(l, "=")
		if len(parts) == 2 {
			keys = append(keys, parts[0])
		}
		output.Exec(l)
	}

	output.Println("Setting environment for %s with [%s]!", name, strings.Join(keys, ", "))

	return nil
}
