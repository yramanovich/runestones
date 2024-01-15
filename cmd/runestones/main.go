package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yramanovich/runestones/contents"
	"github.com/yramanovich/runestones/handlers"
	"github.com/yramanovich/runestones/log"
	"github.com/yramanovich/runestones/repository"
	"github.com/yramanovich/runestones/runestones"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	repo, err := repository.NewPostgres(ctx)
	if err != nil {
		logger.Error("init postgres repository", "err", err)
		return
	}

	router := handlers.SetupHandlers(runestones.NewService(contents.NewFilesystem(""), repo))

	server := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		BaseContext:       func(net.Listener) context.Context { return log.IntoContext(ctx, logger) },
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("listen and serve", "err", err)
		}
	}()
	<-ctx.Done()

	logger.Info("Received interruption signal")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("server shutdown", "err", err)
	}
}
