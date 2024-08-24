package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Yukariin/antizapret-sing-box/geosite_antizapret"
)

func main() {
	var (
		token, owner, repo string
	)

	flag.StringVar(&token, "token", "", "Personal Access Token to create release and upload assets")
	flag.StringVar(&owner, "owner", "", "Repo owner")
	flag.StringVar(&repo, "repo", "", "Repo name")
	flag.Parse()

	if token == "" || owner == "" || repo == "" {
		fmt.Println("You should provide token, owner and repo name!")
		flag.Usage()
		os.Exit(1)
	}

	generator := geosite_antizapret.NewGenerator(geosite_antizapret.WithGitHubClient(token, owner, repo))

	if err := generator.GenerateAndUpload(context.Background()); err != nil {
		log.Fatalf("Failed to generate and upload assets: %v", err)
	}
}
