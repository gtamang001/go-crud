package main

import (
	"log"
	"net/http"

	"github.com/gtamang001/go-crud/controllers"
	"github.com/gtamang001/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8080", nil)
}
