package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/saromanov/born/provider"
	structs "github.com/saromanov/born/structs/v1"
	"github.com/saromanov/gitstar"
)

var errGithubClientNotDefined = errors.New("github client is not defined")

// Options represents settings for setup Github
type Options struct {
	Client *github.Client
}

type client struct {
	client *gitstar.GitStar
}

// New creates init of connection to Github
func New(opt Options) (provider.Provider, error) {
	if opt.Client == nil {
		return nil, errGithubClientNotDefined
	}
	c := gitstar.New(opt.Client)
	return client{
		client: c,
	}, nil
}

// Teams returns list of the teams for Github account
func (c client) Teams(u *structs.User) ([]*structs.Team, error) {
	client := c.client.Client()

	opts := &github.ListOptions{
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

// Repo returns specified repo
func (c client) Repo(u *structs.User, id int64) (*structs.Repo, error) {
	client := c.client.Client()
	repo, _, err := client.Repositories.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return toRepo(repo), nil
}

// GetContent returns content of the file
func (c client) GetContent(u *structs.User, repo, path string)(*structs.ContentFile, error) {
	client := c.client.Client()
	data, _, err := client.Repositories.GetContents(context.Background(), "saromanov", repo, path)
	if err != nil {
		return nil, err
	}
	return toContent(data), nil
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

// toRepo provides converting from github representation
// of repository into Born representation
func toRepo(r *github.Repository) *structs.Repo {
	return &structs.Repo{
		ID:       *r.ID,
		Owner:    fmt.Sprintf("%d", *r.Owner.ID),
		Name:     *r.Name,
		FullName: *r.FullName,
		CloneURL: *r.CloneURL,
	}
}
