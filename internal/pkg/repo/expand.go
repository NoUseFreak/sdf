package repo

import (
	"fmt"
	"regexp"

	"github.com/NoUseFreak/sdf/internal/pkg/utils"
)

type RepoExpander struct {
	shorts map[string]string
}

func NewRepoExpander() *RepoExpander {
	return &RepoExpander{
		shorts: map[string]string{
			"":   "github.com",
			"g":  "github.com",
			"bb": "bitbucket.org",
		},
	}
}

func (rs *RepoExpander) Expand(shortRepo string) string {
	r := regexp.MustCompile("^((?P<short>[a-z]+)/)?(?P<org>[^/]+)/(?P<name>[^/]+)")
	matches := utils.ReSubMatchMap(r, shortRepo)

	if _, ok := rs.shorts[matches["short"]]; ok {
		return fmt.Sprintf(
			"%s/%s/%s",
			rs.shorts[matches["short"]],
			matches["org"],
			matches["name"],
		)
	}

	return shortRepo
}
