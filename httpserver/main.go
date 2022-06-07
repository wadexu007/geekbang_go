package main

import (
	"log"
	"net/http"

	"httpserver/controllers"
	"httpserver/models"

	"github.com/golang/glog"
)

func main() {
	glog.Info("Starting http server")

	roothandler := models.WrapHandlerWithLogging(http.HandlerFunc(controllers.RootHandler))
	healthzhandler := models.WrapHandlerWithLogging(http.HandlerFunc(controllers.Healthz))

	http.Handle("/", roothandler)
	http.Handle("/healthz", healthzhandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
