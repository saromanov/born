package main

import (
	"github.com/saromanov/born/provider"
	"github.com/saromanov/born/provider/github"
	"github.com/urfave/cli"
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
func setupProvider(name string) (provider.Provider, error) {
	github.New(github.Options{})
}
