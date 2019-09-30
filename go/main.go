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
	http.HandleFunc("/upload", uploadWatchInfo)
	http.ListenAndServe(":8080", nil)
}

//This function get watch id and brand using getWatch function and passes them to the gohtml template file
func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	scanWatches, err := db.Query("SELECT * FROM watches order by id")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	// do i need to add ";" at end of below function?
	// check for naming convention in mysql vs code vs frontend

	watch := watchInfo{}
	watchSlice := []watchInfo{}

	for scanWatches.Next() {
		var id int
		var brand string
		err = scanWatches.Scan(&id, &brand)
		if err != nil {
			fmt.Println("sW.Scan didn't work")
		}
		watch.ID = id
		watch.Brand = brand
		watchSlice = append(watchSlice, watch)

	}
	tpl.ExecuteTemplate(w, "pic.gohtml", watchSlice)

}

func uploadWatchInfo(w http.ResponseWriter, r *http.Request) {
	// db := dbConn()
	// defer db.Close()
	// if r.Method == "POST" {
	// 	b := r.FormValue("brand")
	// 	insForm, err := db.Prepare("INSERT INTO watches (brand) VALUES ?")
	// 	if err != nil {
	// 		fmt.Println("insert didn't work")
	// 	}
	// 	insForm.Exec(b)
	// 	log.Println("INSERT: Name: " + b)
	tpl.ExecuteTemplate(w, "upload.gohtml", nil)
}

// defer db.Close()
// }

//this is handling an error and can be called it in other page functions
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println("Index did not work. error in index")
	}
}
