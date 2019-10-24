package env

import (
	"strings"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func Cat(name string) error {
	m, err := readModelFile(modelPath(name))
	if err != nil {
		return err
	}

	for _, l := range strings.Split(m.Data, "\n") {
		output.Println(l)
	}

	return nil
}
