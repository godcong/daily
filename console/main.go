package main

import (
	"github.com/godcong/daily"
)

func main() {
	repos := daily.GetRepos()

	if rrepos, b := (repos).([]interface{}); b {
		for _, repo := range rrepos {
			if rr, b := (repo).(map[string]interface{}); b {
				if v, b := rr["url"]; b {
					if vv, b := v.(string); b {
						repoInfo := daily.GetRepoInfo(vv)
						//log.Println(v)
						daily.GetRepoUpdateAt(repoInfo)
					}

				}
			}

		}
	}

}
