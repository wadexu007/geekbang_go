package controllers

import "net/http"

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[INFO] Ping ok!\n"))
}
