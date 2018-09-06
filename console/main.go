package main

import (
	"github.com/godcong/daily"
	"log"
)

func main() {
	repos := daily.GetRepos()

	if rrepos, b := (repos).([]interface{}); b {
		for _, repo := range rrepos {
			if rr, b := (repo).(map[string]interface{}); b {
				if v, b := rr["url"]; b {
					if vv, b := v.(string); b {
						data := daily.GetRepoDate(vv)
						log.Println(data)
					}

				}
			}

		}
	}

}
