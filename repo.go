package daily

import (
	"log"
)

//GetRepos GetRepos
func GetRepos() interface{} {

	repos := HTTPGet(RepoUrl, defaultHeader(), defaultQuery())
	return repos
}

func FilterRepoUrl() {

}

//GetRepoInfo  GetRepoInfo
func GetRepoInfo(url string) interface{} {
	repos := HTTPGet(url, defaultHeader(), defaultQuery())
	return repos
}

//GetRepoUpdateAt GetRepoUpdateAt
func GetRepoUpdateAt(repo interface{}) {
	if repo == nil {
		return
	}
	log.Println(repo)
	if v, b := repo.(map[string]interface{}); b {
		if vv, b := v["updated_at"]; b {
			log.Println(vv)
		}
	}

}
