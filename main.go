package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"

	"swagger/handler"
	"syscall"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product", log.LstdFlags)
	l.Println("helo")

	//:::::RoUTING server

	mu := mux.NewRouter()
	ph := handler.NewProductHandler(l)
	//Creatuign routers
	getMu := mu.Methods(http.MethodGet).Subrouter()
	postMu := mu.Methods(http.MethodPost).Subrouter()
	putMu := mu.Methods(http.MethodPut).Subrouter()

	//ENDPOINTS
	putMu.Use(ph.MiddlewareProductsValidate)
	postMu.Use(ph.MiddlewareProductsValidate)

	opts := middleware.RedocOpts{
		SpecURL: "swagger.yaml",
	}
	sh:= middleware.Redoc(opts,nil)




	getMu.HandleFunc("/", ph.GetProducts)
	getMu.Handle("/docs", sh)
	getMu.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	postMu.HandleFunc("/", ph.PostProducts)
	putMu.HandleFunc("/id:[0-9]+", ph.PutProducts)



	s := http.Server{
		Addr:         ":9000",
		Handler:      mu,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}

	//::::STARTING server

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Println("Error while starting server %s", err)
			os.Exit(1)
		}
	}()

	//:::::SIGNAL FOR shutdown server

	si := make(chan os.Signal)
	signal.Notify(si, os.Interrupt)
	signal.Notify(si, syscall.SIGINT)
	signal.Notify(si, os.Kill)

	sig := <-si

	l.Println("Gotta signal ::: ", sig)

	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := s.Shutdown(timeout)
	if err != nil {
		l.Println("Error while shut down server %s", err)
		os.Exit(1)
	}

}
