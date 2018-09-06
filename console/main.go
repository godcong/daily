package main

import (
	"github.com/godcong/daily"
	"log"
)

func main() {
	result := daily.DefaultClient().HTTPGet("https://api.github.com/users/godcong", nil, nil)
	log.Println(result)
}
