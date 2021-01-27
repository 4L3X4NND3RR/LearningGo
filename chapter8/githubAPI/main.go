package main

import (
	"log"
	"os"

	"github.com/levigross/grequests"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

type Repo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Forks int `json:"forks"`
	Private bool `json:"private"`
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func main() {
	var repos []Repo
	var repoUrl = "https://api.github.com/users/torvalds/repos"
	resp := getStats(repoUrl)
	resp.JSON(&repos)
	log.Println(repos)
}
