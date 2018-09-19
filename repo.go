package daily

import (
	"log"
)

type Repository struct {
	url   string
	repos interface{}
}

type RepositoryInfo interface{}

func (r *Repository) GetRepoInfo(name string) RepositoryInfo {
	return getRepoInfo(r)
}

//GetRepos GetRepos
func Repos(url string, opts ...map[string]string) *Repository {
	repos := HTTPGet(url, defaultHeader(), defaultQuery())
	return repos
}

func toMap(source interface{}) map[string]interface{} {
	if v, b := source.(map[string]interface{}); b {
		return v
	}
	return nil
}

//GetRepoInfo  GetRepoInfo
func getRepoInfo(repos *Repository) interface{} {
	p := toMap(repos.repos)
	if p == nil {
		return nil
	}
	for k, v := range p {
		log.Println(k, v)
	}

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
