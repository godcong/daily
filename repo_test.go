package daily

import "testing"

const RepoUrl = "https://api.github.com/users/godcong/repos"
const RootUrl = "https://api.github.com/users/godcong"

func TestRepos(t *testing.T) {
	Repos(RepoUrl)
}
