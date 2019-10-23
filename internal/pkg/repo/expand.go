package repo

import (
	"errors"
	"regexp"
	"strings"

	"github.com/NoUseFreak/sdf/internal/pkg/utils"
)

type RepoExpander struct {
	shorts     map[string]string
	transports map[string]map[string]string
}

func NewRepoExpander() *RepoExpander {
	return &RepoExpander{
		shorts: map[string]string{
			"":   "github.com",
			"g":  "github.com",
			"bb": "bitbucket.org",
		},
		transports: map[string]map[string]string{
			"github.com": {
				"git":   "git@github.com:[org]/[name].git",
				"https": "https://github.com/[org]/[name].git",
			},
		},
	}
}

func (rs *RepoExpander) Explode(input string) (string, string, string) {
	r := regexp.MustCompile("^(([a-z]+]?://|([a-z]+@))?(?P<short>[^/:]+)[/:])?(?P<org>[^/]+)/(?P<name>[^/]+)")
	matches := utils.ReSubMatchMap(r, input)

	if strings.HasSuffix(matches["name"], ".git") {
		matches["name"] = matches["name"][0 : len(matches["name"])-4]
	}

	return matches["short"], matches["org"], matches["name"]
}

func (rs *RepoExpander) ExpandPlatform(input string) string {
	if plat, ok := rs.shorts[input]; ok {
		return plat
	}
	return input
}

func (rs *RepoExpander) ExpandTransport(transport string, platfrom string, org string, name string) (string, error) {
	platformtpls, ok := rs.transports[platfrom]
	if !ok {
		return "", errors.New("Unknown platform")
	}
	repo, ok := platformtpls[transport]
	if !ok {
		return "", errors.New("Transport not known for platform")
	}

	repo = strings.ReplaceAll(repo, "[plat]", platfrom)
	repo = strings.ReplaceAll(repo, "[org]", org)
	repo = strings.ReplaceAll(repo, "[name]", name)

	return repo, nil
}
