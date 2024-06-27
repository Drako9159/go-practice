package main

import (
	"GoBaby/cmd/web/routes/mainRoute"
	"GoBaby/cmd/web/routes"
	"context"
	"log"
	"log/slog"
	"net/http"
	repository_domain "GoBaby/cmd/web/domain/repository"
	
)

func InitRoutes() {
	routes.LogRender()
	routes.ErrorRender()
	mainRoute.MainRender()
	mainRoute.ClockRender()
}

func main() {
	InitRoutes()
	repository_domain.InitializeDB()

	mux := routes.GetMuxInstance()
	fileServer := routes.GetFileServerInstance()

	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	slog.Log(context.TODO(), slog.LevelInfo, "Starting server on port 4000")
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())
	log.Fatal(err)
}