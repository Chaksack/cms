package main

import (
	"net/http"
	"trainingProject/cms"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/page/", cms.ServePost)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return 
	}
}
