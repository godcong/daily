package daily

import "net/url"

func GetRepos() interface{} {
	values := make(url.Values)
	values.Add("per_page", "100")
	repos := HTTPGet(RepoUrl, nil, values)
	return repos
}

func FilterRepoUrl() {

}

func GetRepoInfo(url string) interface{} {
	repos := HTTPGet(RepoUrl, nil, nil)
	return repos
}

func GetRepoUpdateAt(repo interface{}) {

}
