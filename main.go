package main

import (
	"fmt"
	"github.com/Ariorhx/jiraClone/CannotResolve"
	"github.com/Ariorhx/jiraClone/RestServer"
	"github.com/gorilla/mux"
)

func main() {
	initData()
	createServer()
}

func createServer() {
	r := mux.NewRouter()
	s := RestServer.Server(r)

	r.HandleFunc("/", RestServer.HelloWorld).Methods("GET")
	r.HandleFunc("/Tasks", RestServer.GetTasks).Methods("GET")
	r.HandleFunc("/Tasks", RestServer.AddTask).Methods("POST")
	r.HandleFunc("/Tasks/{taskHeader}", RestServer.EmptySuit).Methods("PATCH")
	r.HandleFunc("/Tasks/{taskHeader}", RestServer.DeleteTask).Methods("DELETE")

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initData() {
	fmt.Println(CannotResolve.GetUsers())
	fmt.Println(CannotResolve.GetPriority())
}
