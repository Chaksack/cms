package cms

import (
	"net/http"
	"time"
	"strings"
)

//ServePage serves a page based on the route matched. This will match any URL
//beginning with /page
func ServePage(w http.ResponseWriter, r *http.Request)  {
	path := strings.TrimLeft(r.URL.Path, "/page/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		Title: strings.ToTitle(path),
		Content: "Here is my page",
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}

//ServePost serves a post
func ServePost(w http.ResponseWriter, r *http.Request)  {
	path := strings.TrimLeft(r.URL.Path, "/post/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		Title: strings.ToTitle(path),
		Content: "Here is my post",
	}
	Tmpl.ExecuteTemplate(w, "post", p)
}

//Handlenew handles preview logic
func HandleNew(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title: title,
				Content: content,
			})
			return
		}
		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title: title,
				Content: content,
			})
			return
		}
	default:
	http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Go Projects CMS",
		Content: "Welcome to out home page!",
		Posts: []*Post{
			&Post{
				Title: "Hello, World!",
				Content: "Hello world! Thanks for coming to the site.",
				DatePublished: time.Now(),
			},
			&Post{
				Title: "A Post with Comments",
				Content: "Here's a controversial post. It's sure to attract comments.",
				DatePublished: time.Now() .Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author: "Ben Tranter",
						Comment: "Nevermind, I guess I just commented on my post...",
						DatePublished: time.Now() .Add(-time.Hour /2),
					},
				},
			},
		},
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}
