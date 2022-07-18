//https://ieftimov.com/posts/testing-in-go-testing-http-servers/
//https://blog.csdn.net/aCfeng/article/details/122162272

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/logutils"
	"main.go/config"
	"main.go/controllers"
)

func main() {

	//slightly better logging in Go from hashicorp
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(config.Conf.LOG_LEVEL),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	log.Println("[INFO] Red Configuration")
	log.Println("[INFO] " + config.Conf.FILE_PATH)
	log.Println("[INFO] " + config.Conf.LOG_LEVEL)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/healthz", controllers.HealthzHandler).Methods("GET")
	r.HandleFunc("/pizzas", controllers.GetPizzas).Methods("GET")
	r.HandleFunc("/pizzas", controllers.UpdatePizzas).Methods("POST")
	r.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	r.HandleFunc("/orders", controllers.PlaceOrders).Methods("POST")
	r.HandleFunc("/orders/{id}", controllers.GetOrderByPizzaID).Methods("GET")

	// Bind to a port and pass our router in
	log.Println("[INFO] Start http server")
	log.Fatal(http.ListenAndServe(":8080", r))

	// Listen on application shutdown signals.
	listener := make(chan os.Signal, 1)
	signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
	log.Println("[INFO] http server: received a shutdown signal:", <-listener)

}
