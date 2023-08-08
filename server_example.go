package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"time"
)

var db *gorm.DB

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Student struct {
	gorm.Model
	Name      string
	Birthdate time.Time
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/todo.html.tmp"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func StudentIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/students/index.html"))
	var students []Student
	db.Find(&students)
	fmt.Println(students)
	tmpl.Execute(w, students)
}

func main() {
	dsn := "host=localhost user=davidzabner dbname=gormexample port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("%s\n", err)
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Student{})
	var students []Student
	db.Find(&students)
	fmt.Println(students)
	http.HandleFunc("/", TodoHandler)
	http.HandleFunc("/students/", StudentIndex)
	http.ListenAndServe(":80", nil)
}
