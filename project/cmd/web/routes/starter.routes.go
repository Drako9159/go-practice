package routes

import (
	"net/http"
	"io/fs"
	"GoBaby/ui"
)


var (
	mux = http.NewServeMux()
	fileServer http.Handler
)

func GetMuxInstance() *http.ServeMux {
	return mux
}

func GetFileServerInstance() http.Handler {

	if fileServer == nil {
		staticFS, _ := fs.Sub(ui.Content, "static")
		fileServer = http.FileServerFS(staticFS)
	}

	return fileServer

}