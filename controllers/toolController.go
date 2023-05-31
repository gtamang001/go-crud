package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gtamang001/go-crud/initializers"
	"github.com/gtamang001/go-crud/models"
)

var tmpl = template.Must(template.ParseGlob("views/*"))

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	db := initializers.DbConn()
	selDB, err := db.Query("SELECT * FROM tools ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tool := models.Tool{}
	res := []models.Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(id) + " | name " + name + " | category " + category + " | url " + url + " | rating " + string(rating) + " | notes " + notes)

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
		res = append(res, tool)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//Show handler
func Show(w http.ResponseWriter, r *http.Request) {
	db := initializers.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := models.Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}

		log.Println("Showing Row: Id " + string(id) + " | name " + name + " | category " + category + " | url " + url + " | rating " + string(rating) + " | notes " + notes)

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
	}
	tmpl.ExecuteTemplate(w, "Show", tool)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := initializers.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := models.Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
	}

	tmpl.ExecuteTemplate(w, "Edit", tool)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := initializers.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		category := r.FormValue("category")
		url := r.FormValue("url")
		rating := r.FormValue("rating")
		notes := r.FormValue("notes")
		insForm, err := db.Prepare("INSERT INTO tools (name, category, url, rating, notes) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, category, url, rating, notes)
		log.Println("Insert Data: name " + name + " | category " + category + " | url " + url + " | rating " + rating + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := initializers.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		category := r.FormValue("category")
		url := r.FormValue("url")
		rating := r.FormValue("rating")
		notes := r.FormValue("notes")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE tools SET name=?, category=?, url=?, rating=?, notes=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, category, url, rating, notes, id)
		log.Println("UPDATE Data: name " + name + " | category " + category + " | url " + url + " | rating " + rating + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	db := initializers.DbConn()
// 	tool := r.URL.Query().Get("id")
// 	fmt.Println("printing the value of tool as: ", tool)
// 	delForm, err := db.Prepare("DELETE FROM tools WHERE id=?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	delForm.Exec(tool)
// 	log.Println("DELETE " + tool)
// 	defer db.Close()
// 	http.Redirect(w, r, "/", 301)
// }

// added methods for delete
func Delete(w http.ResponseWriter, r *http.Request) {
	// Retrieve the item ID from the request, assuming it's passed in the query string
	fmt.Println("Calling Delete hangle")
	db := initializers.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	fmt.Println("NID value is ", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := models.Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
	}

	tmpl.ExecuteTemplate(w, "Delete", tool)

}

func Confirm(w http.ResponseWriter, r *http.Request) {
	// Handle the delete confirmation form submission here
	if r.Method == http.MethodPost {
		// Retrieve the item ID from the form data
		itemID := r.FormValue("id")

		// Perform the delete operation or further processing based on the item ID
		fmt.Printf("Deleting item with ID: %s\n", itemID)
		db := initializers.DbConn()
		delForm, err := db.Prepare("DELETE FROM tools WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(itemID)
		log.Println("DELETE " + itemID)
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}
