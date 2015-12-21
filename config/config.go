package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	filepath "path/filepath"
	"regexp"

	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/spec"
	apiclient "github.com/jawspeak/go-stash-restclient/client"
	"golang.org/x/crypto/ssh/terminal"
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

func ParseJsonFileStripComments(path string, conf interface{}) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("read file: ", path, " working dir: ", wd, " absolute path: ", abspath)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		panic(err)
	}
	commentStripper := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	file = commentStripper.ReplaceAll(file, nil)
	json.Unmarshal(file, &conf)
}

func SetupConfig() {
	var conf goStashRestClientConfig
	ParseJsonFileStripComments("./config.json", &conf)
	validateRequiredField("host", &conf.Host)
	validateRequiredField("username", &conf.Username)
	if &conf.Password == nil || len(conf.Password) == 0 {
		fmt.Printf("Enter your password for %s: ", conf.Username)
		bytePassword, err := terminal.ReadPassword(0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println()
		conf.Password = string(bytePassword)
	}
	validateRequiredField("password", &conf.Password)

	doc, err := spec.New(apiclient.SwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	transport := httptransport.New(doc)
	transport.Host = conf.Host
	fmt.Println("using host", conf.Host)
	// Assumes basic auth. TODO enable the config.json to take different mechanisms, OR integrate with swagger spec file what it says is supported.
	transport.DefaultAuthentication = httptransport.BasicAuth(conf.Username, conf.Password)
	apiclient.Default.SetTransport(transport)
}
