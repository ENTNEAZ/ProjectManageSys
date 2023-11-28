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

	http.HandleFunc("/api/getAllOrSpecifiedResearchRoomWorker", apiHandler.GetAllOrSpecifiedResearchRoomWorker)
	http.HandleFunc("/api/addOrUpdateResearchRoomWorker", apiHandler.AddOrUpdateResearchRoomWorker)
	http.HandleFunc("/api/deleteResearchRoomWorker", apiHandler.DeleteResearchRoomWorker)

	http.HandleFunc("/api/getAllOrSpecifiedProject", apiHandler.GetAllOrSpecifiedProject)
	http.HandleFunc("/api/addOrUpdateOrDeleteProject", apiHandler.AddOrUpdateOrDeleteProject)
	http.HandleFunc("/api/findAllWorkerInProject", apiHandler.FindAllWorkerInProject)
	http.HandleFunc("/api/addOrDeleteProjectWorker", apiHandler.AddOrDeleteProjectWorker)

	http.HandleFunc("/api/findAllParticipantInProject", apiHandler.FindAllParticipantInProject)
	http.HandleFunc("/api/addOrDeleteProjectParticipant", apiHandler.AddOrDeleteProjectParticipant)

	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/", fs)

	http.ListenAndServe(":56789", nil)
}
