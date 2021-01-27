package main

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

// Response of API
type Response struct {
	License struct {
		Name string `json:"name"`	
		Description string `json:"description"`
	} `json:"license"`
}

func main() {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://api.github.com/graphql")

	// make a request to GitHub API
	req := graphql.NewRequest(`query {
		license(key: "apache-2.0") {
			name
			description
		}
	}`)
	var GitHubToken = os.Getenv("GITHUB_TOKEN")
	req.Header.Add("Authorization", "bearer " + GitHubToken)

	// define a Context for the request
	ctx := context.Background()
	
       // run it and capture the response
	var respData Response
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData.License.Description)
}
