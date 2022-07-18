//https://ieftimov.com/posts/testing-in-go-testing-http-servers/
//https://blog.csdn.net/aCfeng/article/details/122162272

package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"main.go/config"
	"main.go/controllers"
)

func main() {

	log.Println("Red Configuration")
	// configuration := config.GetConfig()
	config.FILE_PATH = config.Conf.FILE_PATH
	log.Println(config.Conf.FILE_PATH)
	log.Println(config.Conf.DB_HOST)
	log.Println(config.Conf.DB_NAME)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/healthz", controllers.HealthzHandler).Methods("GET")
	r.HandleFunc("/pizzas", controllers.GetPizzas).Methods("GET")
	r.HandleFunc("/pizzas", controllers.UpdatePizzas).Methods("POST")
	r.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	r.HandleFunc("/orders", controllers.PlaceOrders).Methods("POST")
	r.HandleFunc("/orders/{id}", controllers.GetOrderByID).Methods("GET")

	// Bind to a port and pass our router in
	log.Println("Start http server")
	log.Fatal(http.ListenAndServe(":80", r))
}
