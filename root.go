package daily

func GetRoot() interface{} {
	repos := HTTPGet(RootUrl, nil, nil)
	return repos
}
