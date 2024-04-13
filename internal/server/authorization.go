package server

import "net/http"

const (
	apiGoogleKey = ""
	apiGitHubKey = ""
)

func Google(w http.ResponseWriter, r *http.Request) {}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {}

func GitHub(w http.ResponseWriter, r *http.Request) {}

func GitHubCallback(w http.ResponseWriter, r *http.Request) {}
