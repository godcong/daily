package daily

import (
	"log"
)

type Repository struct {
	repos interface{}
}

//GetRepos GetRepos
func GetRepos() interface{} {
	repos := HTTPGet(RepoUrl, defaultHeader(), defaultQuery())
	return repos
}

//GetRepoInfo  GetRepoInfo
func GetRepoInfo(repos interface{}) interface{} {

}

//GetRepoUpdateAt GetRepoUpdateAt
func GetRepoUpdateAt(repo interface{}) {
	if repo == nil {
		return
	}
	//log.Println(repo)
	if v, b := repo.(map[string]interface{}); b {
		if vv, b := v["updated_at"]; b {
			log.Println(vv)
		}
	}

}
