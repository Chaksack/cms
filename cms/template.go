package cms

import (
	"html/template"
	"time"
)

//Tmpl is an exported variable
var Tmpl = template.Must(template.ParseGlob("../template/*"))

type Page struct {
	Title string
	Content string
	Posts []*Post
}

type Post struct {
	Title string
	Content string
	DatePublished time.Time
	Comments []*Comment
}

type Comment struct {
	Author string
	Comment string
	DatePublished time.Time
	Comments []*Comment
}

