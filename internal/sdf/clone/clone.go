package clone

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"

	"github.com/NoUseFreak/sdf/internal/pkg/repo"
)

func CloneRepo(repoInput string) (string, error) {
	r := repo.NewFromString(repoInput)

	fullRepo := r.URL
	logrus.Debugf("Expanded %s to %s", repoInput, fullRepo)

	targetDir := path.Join(
		viper.GetString("projectdir"),
		r.TargetDir(),
	)
	logrus.Debugf("Cloning into %s", targetDir)

	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:        fullRepo,
		Progress:   os.Stderr,
		RemoteName: "origin",
	})

	if err != nil {
		switch err {
		case git.ErrRepositoryAlreadyExists:
			logrus.Info("Repo already exists")
		default:
			return "", err
		}
	}

	return targetDir, nil
}
