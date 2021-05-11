package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var fieldID = "customfield_10038"

	mapping, response, err := atlassian.Issue.Field.Context.ProjectsContext(context.Background(), fieldID, nil, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Code", response.StatusCode)
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(mapping.Total)

	fmt.Println(string(response.BodyAsBytes))
}
