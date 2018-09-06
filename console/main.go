package main

import (
	"github.com/godcong/daily"
	"log"
)

func main() {
	//result := daily.DefaultClient().HTTPGet("https://api.github.com/users/godcong", nil, nil)
	//log.Println(result)

	repos := daily.GetRepos()
	log.Println(repos)

	if rrepos, b := (repos).([]interface{}); b {
		for _, repo := range rrepos {
			if rr, b := (repo).(map[string]interface{}); b {
				if v, b := rr["url"]; b {
					log.Println(v)
				}
			}

		}
	}

}
