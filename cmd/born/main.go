package main

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/saromanov/born/provider"
	"github.com/saromanov/born/provider/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "GITHUB_TOKEN",
		Name:   "github-token",
		Usage:  "using of github token for test",
	},
	cli.StringFlag{
		EnvVar: "BORN_SERVER",
		Name:   "born-server",
		Usage:  "server of the born",
	},
}

// setupProvider provides initialization of provider
// at this moment, its support only Github
func setupProvider(c *cli.Context) (provider.Provider, error) {
	token := c.String("github-token")
	ctx := context.Background()
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, sts)
	client := github.NewClient(tc)
	return github.New(github.Options{
		Client: client
	})
}
