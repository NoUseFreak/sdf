package env

import (
	"io/ioutil"
	"os"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func Save(name, tmpFile string) error {
	b, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	m := model{
		Name: name,
		Data: string(b),
	}

	if err := writeModelFile(modelPath(name), m); err != nil {
		return err
	}

	output.Println("Saved %s!", name)
	return nil
}
