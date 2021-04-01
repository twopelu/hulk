package lib

import (
	"fmt"
	"testing"

	"github.com/ktrysmt/go-bitbucket"
)

func Test(t *testing.T) {
	c := bitbucket.NewBasicAuth("twopelu@hotmail.com", "Shinobi*3FBBE")

	fmt.Println("List repos...")
	optr := &bitbucket.RepositoriesOptions{
		Owner: "twopelu",
		Role:  "",
	}

	repos, err1 := c.Repositories.ListForAccount(optr)
	if err1 != nil {
		fmt.Println("Error listing repos:", err1)
		panic(err1)
	}
	fmt.Println("List number of repos:", repos.Size)

	for _, repo := range repos.Items {
		fmt.Println("List repo with name:", repo.Name)
	}

	repong6 := repos.Items[4]

	fmt.Println("List files in repo:", repong6.Name, "...")
	optrf := &bitbucket.RepositoryFilesOptions{
		Owner:    "twopelu",
		RepoSlug: "angular6-project",
		Ref:      "6df9d13dbd815d754172f05b35c113d1c7fb6796",
		Path:     ".gitignore",
	}

	files, err2 := repong6.ListFiles(optrf)
	if err2 != nil {
		fmt.Println("Error listing files:", err2)
		panic(err2)
	}
	fmt.Println("List number of files:", len(files))

	for _, file := range files {
		fmt.Println("List file with name:", file.Path)
	}
}
