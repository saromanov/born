package main

import (
	"context"
	"fmt"
	"os"
	"errors"

	githubl "github.com/google/go-github/github"
	"github.com/saromanov/born/provider"
	"github.com/saromanov/born/provider/github"
	"github.com/saromanov/born/server"
	"github.com/saromanov/born/store"
	"github.com/saromanov/born/store/postgresql"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var errDBCredsIsNotDefined = errors.New("db: username and password is not defined")

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
	cli.StringFlag{
		EnvVar: "DB_USERNAME",
		Name:   "db-username",
		Usage:  "databse username",
	},
	cli.StringFlag{
		EnvVar: "DB_PASSWORD",
		Name:   "db-password",
		Usage:  "databse password",
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
	client := githubl.NewClient(tc)
	return github.New(github.Options{
		Client: client,
	})
}

// setupStore provides initialization of db store
// at this moment its support Postgesql
func setupStore(c *cli.Context) (store.Store, error) {
	username := c.String("db-username")
	password := c.String("db-password")
	if username == "" && password == "" {
		return nil, errDBCredsIsNotDefined
	}
	client, err := postgresql.New(&store.Options{
		Username: c.String("db-username"),
		Password: c.String("db-password"),
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

//setupServer provides create of the new server
func setupServer(p provider.Provider, s store.Store, c *cli.Context) {
	server.Setup(p, s, c.String("born-server"))
}

func run(c *cli.Context) {
	provider, err := setupProvider(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(provider)
	store, err := setupStore(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(store)

	setupServer(provider, store, c)
}

func main() {
	app := cli.NewApp()
	app.Name = "born"
	app.Flags = flags
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
