package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/spec"
	apiclient "github.com/jawspeak/go-stash-restclient/client"
)

type goStashRestClientConfig struct {
	Host     string `json:host`
	Username string `json:username`
	Password string `json:password`
}

func validateRequiredField(field string, configValue *string) {
	if configValue == nil || len(*configValue) == 0 {
		fmt.Println("Required field unset in config.json: ", field)
		os.Exit(1)
	}
}

func SetupConfig() {
	file, err := ioutil.ReadFile("./config.json")
	commentStripper := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	file = commentStripper.ReplaceAll(file, nil)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		panic(err)
	}
	var config goStashRestClientConfig
	json.Unmarshal(file, &config)
	validateRequiredField("host", &config.Host)
	//validateRequiredField("username", config.Username)
	//validateRequiredField("password", config.Password)

	doc, err := spec.New(apiclient.SwaggerJSON, "")
	if err != nil {
		panic(err)
	}
	transport := httptransport.New(doc)
	transport.Host = config.Host
	apiclient.Default.SetTransport(transport)
}
