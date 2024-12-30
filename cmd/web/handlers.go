package main

import (
	"net/http"
	"text/template"
)

func (app *app) homeHandler(w http.ResponseWriter, r *http.Request) {

	posts, err := app.posts.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	temp, err := template.ParseFiles("./assets/templates/home.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	temp.Execute(w, map[string]any{"Posts": posts})
}

func (app *app) createPostHandler(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("./assets/templates/post.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	temp.Execute(w, nil)
}

func (app *app) storePostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = app.posts.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
	)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w,r,"/",http.StatusFound)
}
