package lib

import(
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
)

// Connect to BitBucket and list all the repos for the account,
// user is the email, password in plain text and owner is the username
func ConnectAndListRepos(user string, pass string, owner string) {
	
	c := bitbucket.NewBasicAuth(user, pass)

	opt := &bitbucket.RepositoriesOptions{
		Owner: owner,
	}

	repos, err := c.Repositories.ListForAccount(opt)
	if err != nil {
		fmt.Println("Error listing repos:", err)
		return
	}
	fmt.Println("List number of repos:", repos.Size)

	for _, repo := range repos.Items {
		fmt.Println("List repo with name:", repo.Name)
	}
}