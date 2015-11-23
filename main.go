package main

import (
	"fmt"

	"log"

	apiclient "github.com/jawspeak/go-stash-restclient/client"
	"github.com/jawspeak/go-stash-restclient/client/operations"
)

func main() {
	SetupConfig()

	resp, err := apiclient.Default.Operations.GetCommits(operations.GetCommitsParams{Project: "GO", Repo: "square"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
