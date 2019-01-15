package main

import (
	"github.com/saromanov/born/provider"
	"github.com/saromanov/born/provider/github"
)

// setupProvider provides initialization of provider
// at this moment, its support only Github
func setupProvider(name string) (provider.Provider, error) {
	github.New(github.Options{})
}
