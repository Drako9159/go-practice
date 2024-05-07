package main

import (
	"GoBaby/cmd/web/routes/mainRoute"
	"GoBaby/cmd/web/routes"
	"context"
	"log"
	"log/slog"
	"net/http"
)

func InitRoutes() {
	mainRoute.MainRender()
	mainRoute.ClockRender()
}

func main() {
	InitRoutes()

	mux := routes.GetMuxInstance()
	fileServer := routes.GetFileServerInstance()

	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	slog.Log(context.TODO(), slog.LevelInfo, "Starting server on port 4000")
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())
	log.Fatal(err)
}