package cd

import (
	"io/ioutil"
	"path"

	"github.com/sahilm/fuzzy"
	"github.com/spf13/viper"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func Find(name string) (string, error) {
	rootDir := viper.GetString("projectdir")
	projects := getProjects(rootDir)

	results := fuzzy.Find(name, projects)
	if len(results) == 1 {
		return path.Join(rootDir, results[0].Str), nil
	}
	switch len(results) {
	case 0:
		output.Println("No projects found matching the filter")
	case 1:
		return results[0].Str, nil
	default:
		output.Println("More than 1 match found")
		for _, match := range results {
			output.Println(" - %s", match.Str)
		}
	}
	return "", nil
}

func getProjects(rootDir string) []string {
	dirs := []string{}

	for _, platform := range readDir(rootDir) {
		platformDir := path.Join(rootDir, platform)
		for _, organization := range readDir(platformDir) {
			organizationDir := path.Join(platformDir, organization)
			for _, project := range readDir(organizationDir) {
				dirs = append(dirs, path.Join(platform, organization, project))
			}
		}
	}

	return dirs
}

func readDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		output.Error(err.Error())
	}

	subdirs := []string{}
	for _, f := range files {
		subdirs = append(subdirs, f.Name())
	}

	return subdirs
}
