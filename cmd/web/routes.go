package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.homeHandler)
	mux.HandleFunc("GET /posts/create", app.createPostHandler)
	mux.HandleFunc("POST /posts/create", app.storePostHandler)
	return mux
}
