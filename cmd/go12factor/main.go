package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logrus.Info("Hello World!")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		return
	}
}
