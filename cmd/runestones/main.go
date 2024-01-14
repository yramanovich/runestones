package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/yramanovich/runestones/contents"
	"github.com/yramanovich/runestones/handlers"
	"github.com/yramanovich/runestones/log"
	"github.com/yramanovich/runestones/repository"
	"github.com/yramanovich/runestones/runestones"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.Background()

	fs := contents.NewFilesystem("")
	repo, err := repository.NewPostgres(ctx)
	if err != nil {
		panic(err)
	}

	m := runestones.NewManager(fs, repo)

	router := handlers.SetupHandlers(m)

	server := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		BaseContext:       func(net.Listener) context.Context { return log.IntoContext(ctx, logger) },
	}

	server.ListenAndServe()
}
