package repo

import "path"

type sfdRepo struct {
	URL      string
	Platform string
	Org      string
	Name     string
}

func NewFromString(input string) *sfdRepo {
	expander := NewRepoExpander()

	r := new(sfdRepo)
	r.URL = input

	plat, org, name := expander.Explode(input)
	r.Platform = expander.ExpandPlatform(plat)
	r.Org = org
	r.Name = name
	transportURL, err := expander.ExpandTransport("git", r.Platform, r.Org, r.Name)
	if err == nil {
		r.URL = transportURL
	}

	return r
}

func (r *sfdRepo) TargetDir() string {
	return path.Join(r.Platform, r.Org, r.Name)
}
