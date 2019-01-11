package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/saromanov/born/provider"
	structs "github.com/saromanov/born/structs/v1"
	"github.com/saromanov/gitstar"
)

// Options represents settings for setup Github
type Options struct {
	Client *github.Client
}

type client struct {
	client *gitstar.GitStar
}

// New creates init of connection to Github
func New(opt Options) provider.Provider {
	c := gitstar.New(opt.Client)
	return client{
		client: c,
	}
}

// Teams returns list of the teams for Github account
func (c *client) Teams(u *structs.User) ([]*structs.Team, error) {
	client := c.client.Client()

	opts := github.ListOptions{
		Page: 1,
	}

	var teams []*structs.Team
	for opts.Page > 0 {
		l, resp, err := client.Organizations.List(context.Background(), "", opts)
		if err != nil {
			return nil, err
		}
		teams = append(teams, toTeamList(l)...)
		opts.Page = resp.NextPage
	}
	return teams, nil
}

// toTeamList provides converting from github representation
// of the Team into Born representation of the Team
func toTeamList(lst []*github.Organization) []*structs.Team {
	teams := make([]*structs.Team, len(lst))
	for i := range lst {
		teams[i] = &structs.Team{
			Name:    *lst[i].Name,
			Company: *lst[i].Company,
			ID:      fmt.Sprintf("%d", *lst[i].ID),
		}
	}
	return teams
}
