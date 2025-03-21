package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RonaldAG/go-microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.Default().Flags())
	ph := handlers.NewProducts(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gh)
	
	
	server := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120 *time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}
	go func() {
		err :=	server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}