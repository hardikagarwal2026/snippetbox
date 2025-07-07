package main

import (
	"fmt"
	"net/http"
	"snippetbox/pkg/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w) // Use the notFound helper
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.ServerError(w, err) // Use the serverError helper
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.ServerError(w, err) // Use the serverError helper
	// 	return
	// }

	// w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w) // Use the notFound helper
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.NotFound(w)
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%v", s)
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError helper

		title := "hardik"
		content := "the great"
		expires := "9"

		id, err := app.snippets.Insert(title, content, expires)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

		return
	}

	w.Write([]byte("Create a new snippet..."))
}
