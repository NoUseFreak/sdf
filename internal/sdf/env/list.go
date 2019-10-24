package env

import (
	"io/ioutil"
	"os"
	"path"
)

func List() []model {
	dir := envDir()
	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		return nil
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}

	var models []model
	for _, f := range files {
		if m, err := readModelFile(path.Join(dir, f.Name())); err == nil {
			models = append(models, m)
		}
	}

	return models
}
