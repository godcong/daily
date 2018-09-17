package daily

import (
	"log"
)

type Repository struct {
	url   string
	repos interface{}
}

//GetRepos GetRepos
func Repos(url string, opts ...map[string]string) *Repository {
	repos := HTTPGet(url, defaultHeader(), defaultQuery())
	return repos
}

//GetRepoInfo  GetRepoInfo
func getRepoInfo(repos *Repository) interface{} {
	return nil
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
