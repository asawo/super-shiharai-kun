package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/asawo/api/config"
	"github.com/asawo/api/http"
	"github.com/asawo/api/service"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	exitOK int = iota
	exitError
)

func main() {
	os.Exit(realMain(os.Args))
}

// realMain is start point of service. We initialize required clients
// and logger and pass it to the server struct.
func realMain(args []string) int {
	// Read env vars with config pkg
	env, err := config.ReadFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to read env vars: %s\n", err)
		return exitError
	}

	// Initialize logger
	logger, err := setupLogger(env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup logger: %s\n", err)
		return exitError
	}

	ctx := context.Background()

	// Initialize DB
	db, err := setupDB(ctx, env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup a repository: %s\n", err)
		return exitError
	}

	// Initialize a http server
	service := service.NewService(logger, env, db)
	httpServer, err := http.New(ctx, logger, service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup an HTTP server: %s\n", err)
		return exitError
	}

	httpAddr := fmt.Sprintf(":%d", env.HTTPPort)
	logger.Info("http server listening", zap.String("address", httpAddr))
	httpLn, err := net.Listen("tcp", httpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to listen HTTP port: %d\n", env.HTTPPort)
		return exitError
	}

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return httpServer.Serve(httpLn) })

	// Waiting for SIGTERM or Interrupt signal. If server receives them,
	// gRPC server and http server will shutdown gracefully.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	select {
	case <-sigCh:
		logger.Info("received SIGTERM, exiting server gracefully")
	case <-ctx.Done():
	}

	// Gracefully shutdown server
	httpServer.GracefulStop(ctx)
	if err := wg.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Unhandled errors received: %s\n", err)
		return exitError
	}

	return exitOK
}
