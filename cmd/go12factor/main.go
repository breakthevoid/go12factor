package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.Info("Hello World!")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	serv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	go func() {
		err := serv.ListenAndServe()
		if err != nil {

		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	serv.Shutdown(timeout)
}
