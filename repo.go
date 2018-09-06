package daily

func GetRepos(url string) map[string]interface{} {
	repos := HTTPGet(url, nil, nil)
	return repos
}
