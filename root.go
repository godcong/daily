package daily

//GetRoot GetRoot
func GetRoot() interface{} {
	repos := HTTPGet(RootUrl, nil, nil)
	return repos
}
