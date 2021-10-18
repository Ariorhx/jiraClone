package RestServer

import (
	"encoding/json"
	"fmt"
	"github.com/Ariorhx/jiraClone/CannotResolve"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func EmptySuit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("empty suit")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "%s\n", "Hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println(r.Header)
	text, _ := ioutil.ReadAll(r.Body)
	fmt.Println(text)
}

func GetBody(w http.ResponseWriter, r *http.Request) {
	bodyText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Cannot read text from body\n%v\n", err)
	}
	fmt.Printf("%s", bodyText)
	task := CannotResolve.Task{}
	err = json.Unmarshal([]byte(bodyText), &task)
	if err != nil {
		fmt.Printf("Cannot parse input body to json\n%v\n", err)
	}
	fmt.Printf("%v", task)

}

func AddTask(w http.ResponseWriter, r *http.Request) {
	bodyText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Cannot read text from body\n%v\n", err)
	}
	task := CannotResolve.Task{}
	err = json.Unmarshal([]byte(bodyText), &task)
	if err != nil {
		fmt.Printf("Cannot parse input body to json\n%v\n", err)
	}
	if _, err = w.Write([]byte(CannotResolve.GetDatabase().AddTask(&task))); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["taskHeader"]
	if _, err := w.Write([]byte(CannotResolve.GetDatabase().DeleteTask(header))); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := CannotResolve.GetDatabase().GetTasks()

	for _, task := range tasks {
		task.ValueFromID()
	}

	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Cannot parse task to json\n%v\n", err)
	}
	if _, err = w.Write([]byte(jsonTasks)); err != nil {
		fmt.Printf("%v\n", err)
	}
}
