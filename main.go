package main

import (
	"encoding/json"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func main(){
	tasks = append(tasks, Task{ ID: 1, Title: "Pick up laundry", Description: "", IsImportant: true, CreatedTime: time.Now(), DueTime: time.Now().AddDate(0, 0, 1)})
	tasks = append(tasks, Task{ ID: 2, Title: "Buy milk", Description: "Buy milk from Kroger", CreatedTime: time.Now()})
	tasks = append(tasks, Task{ ID: 3, Title: "Prepare for demo", Description: "JsConf demo", IsImportant: true, IsFinished: true, CreatedTime: time.Now().AddDate(0, 0, -7), DueTime: time.Now().AddDate(0, 0, 10)})
	tasks = append(tasks, Task{ ID: 4, Title: "Run 5k", Description: "At NCRB", CreatedTime: time.Now().AddDate(0, -1, 0)})

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", AddTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetTasks(writer http.ResponseWriter, req *http.Request){
	json.NewEncoder(writer).Encode(tasks)
}

func GetTask(writer http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, task := range tasks {
		if strconv.Itoa(task.ID) == params["id"]{
			json.NewEncoder(writer).Encode(task)
		}
	}
}

func AddTask(writer http.ResponseWriter, req *http.Request){
	var task Task
	json.NewDecoder(req.Body).Decode(&task)
	task.ID = 5
	tasks = append(tasks, task)
	json.NewEncoder(writer).Encode(task)
}

func UpdateTask(writer http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, task := range tasks {
		if strconv.Itoa(task.ID) == params["id"]{
			tasks[index] = task			
			json.NewEncoder(writer).Encode(task)
		}
	}
}

func DeleteTask(writer http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, task := range tasks {
		if strconv.Itoa(task.ID) == params["id"]{
			tasks = append(tasks[:index], tasks[index+1:]...)			
			break
		}
	}
}

type Task struct {
	ID  int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	IsImportant bool `json:"isImportant,omitempty"`
	IsFinished bool `json:"isFinished,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
	DueTime time.Time `json:"dueTime,omitempty"`
}

var tasks []Task