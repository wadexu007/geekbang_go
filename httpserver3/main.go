//https://ieftimov.com/posts/testing-in-go-testing-http-servers/
//https://blog.csdn.net/aCfeng/article/details/122162272

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/config"
)

type Pizza struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Pizzas []Pizza

// var pizzas Pizzas
var file_path string

type Order struct {
	PizzaID  int `json:"pizza_id"`
	Quantity int `json:"quantity"`
	Total    int `json:"total"`
}

type Orders []Order

// var orders Orders

func (ps Pizzas) FindByID(ID int) (Pizza, error) {
	for _, pizza := range ps {
		if pizza.ID == ID {
			return pizza, nil
		}
	}

	return Pizza{}, fmt.Errorf("couldn't find pizza with ID: %d", ID)
}

func UpdatePizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var p Pizza

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}

		// pizzas := append(pizzas, p)
		// json.NewEncoder(w).Encode(pizzas)

		// store pizza data in csv
		record := []string{
			strconv.Itoa(p.ID),
			p.Name,
			strconv.Itoa(p.Price),
		}
		log.Println("Start to write pizza record to csv")
		WriteData(file_path+"pizzas.csv", record)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pizzas := GetPizzasFromCSV(file_path + "pizzas.csv")
	log.Println("get all pizzas")
	json.NewEncoder(w).Encode(pizzas)

}

func GetPizzasFromCSV(fileName string) Pizzas {
	var pizzas Pizzas
	// read data from csv
	records, err := ReadData(fileName)
	if err != nil {
		log.Println("Can't read data from csv")
		return pizzas

	}
	if len(records) == 0 {
		log.Println("Error: No pizzas found")
		return pizzas
	}
	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}
		price, err := strconv.Atoi(record[2])
		if err != nil {
			log.Println(err)
		}

		pizza := Pizza{
			ID:    id,
			Name:  record[1],
			Price: price,
		}
		pizzas = append(pizzas, pizza)
	}
	return pizzas
}

func (orders Orders) GetByID(ID int) (Order, error) {
	for _, order := range orders {
		if order.PizzaID == ID {
			return order, nil
		}
	}

	return Order{}, fmt.Errorf("couldn't find Order with Pizza ID: %d", ID)
}

func PlaceOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pizzas := GetPizzasFromCSV(file_path + "pizzas.csv")

	var o Order

	if len(pizzas) == 0 {
		http.Error(w, "Error: No pizzas found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, "Can't decode body", http.StatusBadRequest)
		return
	}

	p, err := pizzas.FindByID(o.PizzaID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		return
	}

	o.Total = p.Price * o.Quantity
	// orders = append(orders, o)
	// json.NewEncoder(w).Encode(o)

	// store order data in csv
	order_new := []string{
		strconv.Itoa(o.PizzaID),
		strconv.Itoa(o.Quantity),
		strconv.Itoa(o.Total),
	}
	WriteData(file_path+"orders.csv", order_new)
	log.Println("write placed order record to csv succesfully")
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := GetOrdersFromCSV(file_path + "orders.csv")
	if len(orders) == 0 {
		http.Error(w, "No orders found, please palce order", http.StatusNotFound)
		return
	}
	log.Println("get all orders")
	json.NewEncoder(w).Encode(orders)
}

func GetOrdersFromCSV(fileName string) Orders {
	var orders Orders
	// read data from csv
	records, err := ReadData(fileName)
	if err != nil {
		log.Println("Can't read data from csv")
		return orders

	}
	if len(records) == 0 {
		// log.Fatalln("Error: No orders found")
		log.Println("Info: No orders found")
		return orders
	}
	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}
		quantity, err := strconv.Atoi(record[1])
		if err != nil {
			log.Println(err)
		}
		total_price, err := strconv.Atoi(record[2])
		if err != nil {
			log.Println(err)
		}

		order := Order{
			PizzaID:  id,
			Quantity: quantity,
			Total:    total_price,
		}
		orders = append(orders, order)
	}
	return orders
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := GetOrdersFromCSV(file_path + "orders.csv")

	params := mux.Vars(r)
	log.Println("params ====", params)

	if len(orders) == 0 {
		http.Error(w, "Error: No orders found", http.StatusNotFound)
		return
	}

	orderID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	o, err := orders.GetByID(orderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(o)

}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ping ok!\n"))
}

func WriteData(fileName string, record []string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if os.IsNotExist(err) {
		log.Println("file not exist, now try to crete")
		error := os.MkdirAll(file_path, os.ModePerm)
		if error != nil {
			log.Println(error)
		}
		f, err = os.Create(fileName)

	}
	if err != nil {
		log.Println("failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(record); err != nil {
		log.Println("error writing record to file", err)
	}
}

func ReadData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	// if _, err := r.Read(); err != nil {
	// 	return [][]string{}, err
	// }

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func main() {

	log.Println("Red Configuration")
	configuration := config.GetConfig()
	file_path = configuration.FILE_PATH
	log.Println(configuration.FILE_PATH)
	log.Println(configuration.DB_HOST)
	log.Println(configuration.DB_NAME)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/healthz", HealthzHandler).Methods("GET")
	r.HandleFunc("/pizzas", GetPizzas).Methods("GET")
	r.HandleFunc("/pizzas", UpdatePizzas).Methods("POST")
	r.HandleFunc("/orders", GetOrders).Methods("GET")
	r.HandleFunc("/orders", PlaceOrders).Methods("POST")
	r.HandleFunc("/orders/{id}", GetOrderByID).Methods("GET")

	// Bind to a port and pass our router in
	log.Println("Start http server")
	log.Fatal(http.ListenAndServe(":80", r))
}
