package main

import (
	"Database_Homework/apiHandler"
	"net/http"
)

func main() {
	http.HandleFunc("/api/login", apiHandler.Login)
	http.HandleFunc("/api/register", apiHandler.Register)

	http.HandleFunc("/api/getWorkerByID", apiHandler.GetWorkerByID)
	http.HandleFunc("/api/getAllWorker", apiHandler.GetAllWorker)
	http.HandleFunc("/api/addOrUpdateWorker", apiHandler.AddOrUpdateWorker)

	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/", fs)

	http.ListenAndServe(":56785", nil)
}
