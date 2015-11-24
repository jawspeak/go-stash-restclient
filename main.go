package main

import (
	"fmt"

	"log"

	"github.com/davecgh/go-spew/spew"
	apiclient "github.com/jawspeak/go-stash-restclient/client"
	"github.com/jawspeak/go-stash-restclient/client/operations"
	"github.com/jawspeak/go-stash-restclient/config"
)

func main() {
	config.SetupConfig() // CRITICAL!

	// Examples then how to use the API
	_, err := apiclient.Default.Operations.GetPullRequests(operations.GetPullRequestsParams{Project: "<will-404>", Repo: "square"})
	fatalIfErrUnless404(err)
	// more detailed example to handle a 404.
	if err == nil {
		panic("expecting error for 404 not found")
	} else {
		spew.Dump(err)
		if apiErr, ok := err.(operations.APIError); !ok {
			log.Println(".. err not an APIError")
		} else {
			if apiErr.Code != 404 {
				panic("expecting 404")
			}

			if resp404, ok := apiErr.Response.(*operations.GetPullRequestsNotFound); !ok {
				log.Println(".. err response not GetPullRequestsNotFound")
			} else {
				log.Println("this err's response", resp404)
			}
		}
	}

	resp1, err := apiclient.Default.Operations.GetCommits(operations.GetCommitsParams{Project: "GO", Repo: "square"})
	fatalIfErrUnless404(err)
	fmt.Printf("%#v\n\n", resp1.Payload)

	resp2, err := apiclient.Default.Operations.GetPullRequests(operations.GetPullRequestsParams{Project: "GO", Repo: "square"})
	fatalIfErrUnless404(err)
	fmt.Printf("%#v\n\n", resp2.Payload)
	pullRequestId := resp2.Payload.Values[0].ID

	resp3, err := apiclient.Default.Operations.GetPullRequestActivities(
		operations.GetPullRequestActivitiesParams{Project: "GO", Repo: "square", PullRequestID: pullRequestId})
	fatalIfErrUnless404(err)
	fmt.Printf("%#v\n\n", resp3.Payload)
}

func fatalIfErrUnless404(err error) {
	if err != nil {
		if apiErr, ok := err.(operations.APIError); ok {
			if apiErr.Code == 404 {
				log.Print("Not found - skipping", apiErr)
				return
			}
		}
		log.Fatal(err)
	}
}
