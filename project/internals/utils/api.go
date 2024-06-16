package utils

import (
	"fmt"
	"net/http"
	"html/template"
)

func CheckIfPath(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != path {
		fmt.Println("Error: "+r.URL.Path, path)
		http.NotFound(w, r)
		return
	}
}

func IsValidHTTPMethod(method string, acceptedMethod string, w http.ResponseWriter, r *http.Request) bool {
	if method == acceptedMethod {
		w.Header().Set("Allow", acceptedMethod)
		return true
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return false
}

var (
	EmptyFuncMap = template.FuncMap{}
	EmptyStruct = struct{}{}
)
