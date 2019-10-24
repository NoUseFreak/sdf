package env

import (
	"os"
)

func Remove(name string) error {
	return os.Remove(modelPath(name))
}
