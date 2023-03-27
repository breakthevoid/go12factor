package main

import (
	"context"
	"encoding/json"
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
	log := logrus.New()
	log.SetOutput(os.Stdout)

	log.Info("Starting the app..")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("Error happened in JSON marshal: %s", err)
		}
		_, err = writer.Write(jsonResp)
		if err != nil {
			log.Fatal("Error while writing response: %s", err)
			return
		}
	})

	serv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	go func() {
		err := serv.ListenAndServe()
		if err != nil {
			log.Errorf("Server status: %v", err)
			return
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	log.Info("Stopping app..")
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := serv.Shutdown(timeout)
	if err != nil {
		log.Errorf("Error when shutdown app: %v", err)
		return
	}
	log.Info("The app stopped.")
}
