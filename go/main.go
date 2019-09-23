package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type watchInfo struct {
	ID    int
	Brand string
}

var (
	tpl *template.Template
	// cnn, err = sql.Open("mysql", "root:root@tcp(db:3306)/appdb")
)

//function to connect to db
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "appdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+"(db:3306)/"+dbName)
	if err != nil {
		fmt.Println("dbConn not work")
	}
	return db
}

func init() {
	tpl = template.Must(template.ParseGlob("/go/templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

//This function get watch id and brand using getWatch function and passes them to the gohtml template file
func index(w http.ResponseWriter, r *http.Request) {
	watch, err := getWatch("1")
	HandleError(w, err)

	err = tpl.ExecuteTemplate(w, "pic.gohtml", watch)
	HandleError(w, err)

}

//this is handling an error and can be called it in other page functions
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println("Index did not work. error in index")
	}
}

//This function takes id # as input and outputs watchid and corresponding brand

func getWatch(id string) (watch watchInfo, err error) {
	db := dbConn()
	// , err := sql.Open("mysql", "root:root@tcp(db:3306)/appdb")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.QueryRow("SELECT id, brand from watches where id = ?;", id).Scan(&watch.ID, &watch.Brand)
	if err != nil {
		panic(err.Error())
	}

	return watch, err
}
