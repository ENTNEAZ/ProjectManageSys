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
	http.HandleFunc("/api/deleteWorkerByID", apiHandler.DeleteWorkerByID)

	http.HandleFunc("/api/getAllResearchRoom", apiHandler.GetAllResearchRoom)
	http.HandleFunc("/api/addOrUpdateResearchRoom", apiHandler.AddOrUpdateResearchRoom)
	http.HandleFunc("/api/deleteResearchRoom", apiHandler.DeleteResearchRoom)

	http.HandleFunc("/api/getAllOrSpecifiedWorkingArea", apiHandler.GetAllOrSpecifiedWorkingArea)
	http.HandleFunc("/api/addOrUpdateWorkingArea", apiHandler.AddOrUpdateWorkingAreaSubmit)
	http.HandleFunc("/api/addWorkingAreaForResearchRoom", apiHandler.AddWorkingAreaForResearchRoom)
	http.HandleFunc("/api/deleteWorkingAreaForResearchRoom", apiHandler.DeleteWorkingAreaForResearchRoom)

	http.HandleFunc("/api/getAllOrSpecifiedSectary", apiHandler.GetAllOrSpecifiedSectary)
	http.HandleFunc("/api/addOrUpdateSectary", apiHandler.AddOrUpdateSectary)
	http.HandleFunc("/api/deleteSectary", apiHandler.DeleteSectary)

	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/", fs)

	http.ListenAndServe(":56785", nil)
}
