package lib

import (
	"fmt"

	"github.com/ktrysmt/go-bitbucket"
)

// BitBucket Manager
type BitBucketManager struct {
	client *bitbucket.Client
}

// Connect to BitBucket with basic auth
// - user is the email
// - password in plain text
func (bbman *BitBucketManager) Connect(user string, pass string) {
	bbman.client = bitbucket.NewBasicAuth(user, pass)
}

// List all repos of the owner
// - owner is the username
// API Reference:
// https://api.bitbucket.org/2.0/repositories/twopelu
// https://api.bitbucket.org/2.0/repositories/twopelu
func (bbman *BitBucketManager) ListRepos(owner string) ([]bitbucket.Repository, error) {
	fmt.Println("List repos...")
	opt := &bitbucket.RepositoriesOptions{
		Owner: owner,
		Role:  "",
	}

	repos, err := bbman.client.Repositories.ListForAccount(opt)
	if err != nil {
		fmt.Println("Error listing repos:", err)
		return nil, err
	}
	fmt.Println("List number of repos:", repos.Size)

	for _, repo := range repos.Items {
		fmt.Println("List repo with name:", repo.Name)
	}
	return repos.Items, nil
}

// List files in a repo
// - owner is the username
// API Reference:
// https://api.bitbucket.org/2.0/repositories/{username}/{repo_slug}/src/{commit}/{path}
// https://api.bitbucket.org/2.0/repositories/twopelu/angular6-project/src/6df9d13dbd815d754172f05b35c113d1c7fb6796/.gitignore
func (bbman *BitBucketManager) ListFiles(repo *bitbucket.Repository, owner string) ([]bitbucket.RepositoryFile, error) {
	fmt.Println("List files in repo:", repo.Name, "...")
	opt := &bitbucket.RepositoryFilesOptions{
		Owner:    owner,
		RepoSlug: repo.Slug,
		Ref:      "6df9d13dbd815d754172f05b35c113d1c7fb6796",
		Path:     ".gitignore",
	}

	files, err := repo.ListFiles(opt)
	if err != nil {
		fmt.Println("Error listing files:", err)
		return nil, err
	}
	fmt.Println("List number of files:", len(files))

	for _, file := range files {
		fmt.Println("List file with name:", file.Path)
	}
	return files, nil
}
