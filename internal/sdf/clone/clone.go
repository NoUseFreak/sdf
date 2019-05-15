package clone

import (
	"fmt"
	"github.com/NoUseFreak/sdf/internal/pkg/repo"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path"
)

func CloneRepo(repoInput string) (string, error) {
	repoExpander := repo.NewRepoExpander()
	fullRepo := repoExpander.Expand(repoInput)

	targetDir := path.Join(
		viper.GetString("projectdir"),
		fullRepo,
	)

	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:      fmt.Sprintf("git://%s", fullRepo),
		Progress: os.Stderr,
	})

	if err != nil {
		return "", err
	}

	return targetDir, nil
}
