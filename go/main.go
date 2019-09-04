package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

type watch struct {
	id   int
	name string
}

var (
	tpl      *template.Template
	cnn, err = sql.Open("mysql", "root:root@tcp(db:3306)/watchesdb")
)

func init() {
	tpl = template.Must(template.ParseGlob("/go/photosite/templates/*"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM watches WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}

	err := tpl.ExecuteTemplate(w, "pic.gohtml", id)
	HandleError(w, err)

}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println("Index did not work. error in index")
	}
}
