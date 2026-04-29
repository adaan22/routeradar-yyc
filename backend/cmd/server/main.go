package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"routeradar-yyc-backend/internal/api"
	"routeradar-yyc-backend/internal/server"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	corsAllowOrigin := flag.String("cors-allow-origin", "*", "CORS allow origin")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	router := api.NewRouter(*corsAllowOrigin)
	srv := server.New(*addr, router)

	if err := srv.Run(ctx); err != nil {
		// If the server shuts down normally, Shutdown returns http.ErrServerClosed in some cases.
		log.Printf("server stopped: %v", err)
	}
}

