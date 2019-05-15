package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func ExistsIn(repoURL string, devBranch string, mainBranch string) (bool, error) {
	remoteName := "origin"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.NewBranchReferenceName(devBranch),
		RemoteName:    remoteName,
	})
	if err != nil {
		return true, err
	}

	refDev, err := repo.Head()
	if err != nil {
		return true, err
	}

	refMain, err := repo.Reference(plumbing.NewRemoteReferenceName(remoteName, mainBranch), true)
	if err != nil {
		return true, err
	}

	found := false

	cIter, err := repo.Log(&git.LogOptions{From: refMain.Hash()})
	if err != nil {
		return true, err
	}

	cIter.ForEach(func(c *object.Commit) error {
		if refDev.Hash() == c.Hash {
			found = true
		}

		return nil
	})

	return found, nil
}
