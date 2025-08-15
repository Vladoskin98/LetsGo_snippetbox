package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vladoskin98/LetsGo_snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	//Include the navigation partial in the template files.
	/*
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/home.tmpl",
		}

		// Use the template.ParseFiles() function to read the files and store the
		// templates in a template set. Notice that we use ... to pass the contents
		// of the files slice as variadic arguments.
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, r, err) // Use the serverError() helper.
			return
		}

		//Use the ExecuteTemplate() method to write the content of the "base"
		//template as the response body.
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.serverError(w, r, err) // Use the serverError() helper.
		}
	*/
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly! \n\n- Kobayashi Issa"
	expires := 7

	//Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	//Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
