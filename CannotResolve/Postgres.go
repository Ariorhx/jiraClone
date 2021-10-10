package CannotResolve

import (
	"JiraClone/Configuration"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

var db *dataBase
var dOnce sync.Once

type dataBase struct {
	conn *sql.DB
}

type DataBase interface {
	AddTask(t *Task) string
	DeleteTask(header string) string
	GetTasks() []*Task
	GetUsers() map[string]int
	GetPriority() map[string]int
}

func GetDatabase() DataBase {
	dOnce.Do(func() {
		conn, err := sql.Open("postgres", Configuration.ConnStringFromConf())
		if err != nil {
			panic(err)
		}
		db = &dataBase{conn}
	})
	return db
}

func (db *dataBase) AddTask(t *Task) string {
	t.IdFromValue()
	_, err :=
		db.conn.Exec(
			"insert into jira_clone.tasks(\"header\", \"text\", user_id, priority_id, creation_time, end_time) " +
				fmt.Sprintf(
					"values ('%s', '%s', %d, %d, TIMESTAMP '%s', TIMESTAMP '%s')",
					t.Header, t.Text, t.UserId, t.PriorityId, t.CreationTime, t.EndTime))

	if err != nil {
		switch errorText := fmt.Sprintf("%v", err); errorText {
		case "pq: duplicate key value violates unique constraint \"tasks_pkey\"":
			return "Задание с таким зоголовком уже существует"
		default:
			return errorText
		}
	}

	return "Задание успешно создано"
}

func (db *dataBase) DeleteTask(header string) string {
	query := fmt.Sprintf("delete from jira_clone.tasks where \"header\" = '%s'", header)
	_, err := db.conn.Exec(query)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return "Задание удалено"
}

func (db *dataBase) GetTasks() []*Task {
	rows, err := db.conn.Query("select * from jira_clone.tasks")
	if err != nil {
		fmt.Printf("Cannot get tasks from db, because %s\n", err)
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Printf("Cannot close rows, because %s\n", err)
		}
	}()

	var tasks []*Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Header, &task.Text, &task.UserId, &task.PriorityId, &task.CreationTime, &task.EndTime)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, &task)
	}
	return tasks
}

func (db *dataBase) GetUsers() map[string]int {
	return mapFromSQL(db, "select * from jira_clone.users")
}

func (db *dataBase) GetPriority() map[string]int {
	return mapFromSQL(db, "select * from jira_clone.priority")
}

func mapFromSQL(db *dataBase, query string) map[string]int {
	rows, err := db.conn.Query(query)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Printf("Cannot close rows, because %s\n", err)
		}
	}()

	sqlMap := make(map[string]int)
	for rows.Next() {
		var key string
		var val int

		err = rows.Scan(&val, &key)
		if err != nil {
			panic(err)
		}

		sqlMap[key] = val
	}
	return sqlMap
}
