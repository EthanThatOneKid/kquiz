package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"etok.codes/qg/backend/qg/stores/sqlite"
	"etok.codes/qg/backend/server"
	"github.com/diamondburned/listener"
	"github.com/go-chi/chi/v5"
)

var (
	addr       = "localhost:8080"
	sqlitePath = "/tmp/qg.sqlite"
)

func init() {
	flag.StringVar(&addr, "addr", addr, "address to listen on")
	flag.StringVar(&sqlitePath, "sqlite", sqlitePath, "path to SQLite database")
	flag.Parse()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	store, err := sqlite.New(sqlitePath)
	if err != nil {
		log.Fatalln("failed to open SQLite store:", err)
	}
	defer store.Close()

	handler := server.NewHandler(store)
	defer handler.Close()

	r := chi.NewRouter()
	r.Mount("/api/v0", handler)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("listening on", addr)
	if err := listener.HTTPListenAndServeCtx(ctx, &server); err != nil {
		log.Fatalln("server:", err)
	}
}
