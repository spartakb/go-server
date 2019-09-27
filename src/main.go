package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

func init() {

	fmt.Println("Initializing Appplication")
}

func main() {
	addr := ":80"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()
	
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<b>Server Root: v1.003</b>"))
	})
	
	return r
}

func Sum(x int, y int) int {
    return x + y
}
