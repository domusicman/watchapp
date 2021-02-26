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
	dbUser := "dom"
	dbPass := "dom"
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
	// http.HandleFunc("/createwatch", createWatch)
	http.ListenAndServe(":8080", nil)
}

//This function get watch id and brand using getWatch function and passes them to the gohtml template file
func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	sW, err := db.Query("SELECT * FROM watches order by id")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	// do i need to add ";" at end of below function?
	// check for naming convention in mysql vs code vs frontend
	//sW (select watches)

	wa := watchInfo{}
	waS := []watchInfo{}

	for sW.Next() {
		var id int
		var brand string
		err = sW.Scan(&id, &brand)
		if err != nil {
			fmt.Println("sW.Scan didn't work")
		}
		wa.ID = id
		wa.Brand = brand
		waS = append(waS, wa)

	}
	tpl.ExecuteTemplate(w, "pic.gohtml", waS)
	// watch, err := getWatch("1")
	// HandleError(w, err)

	// err = tpl.ExecuteTemplate(w, "pic.gohtml", watch)
	// HandleError(w, err)

}

// func createWatch(w http.ResponseWriter, r *http.Request) {
// 	//open connection to db
// 	//create variables using struct that i'll need. ID and Watch
// 	//create query to input struct elements into db
// 	//create mapped submissions for struct elements
// 	//

// 	//taken from above
// 	var err error

// 	db := dbConn()
// 	defer db.Close()
// 	err = db.Ping()
// 	// insert watch
// 	iw := db.Exec("INSERT INTO watches (ID, Brand) VALUES ($1,$2)", info.ID, info.Brand)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	wa := watchInfo{}
// 	// waS := []watchInfo{}

// 	var id int
// 	var brand string

// 	wa.ID = id
// 	wa.Brand = brand

// 	// tpl.ExecuteTemplate(w, "pic.gohtml", waS)
// 	//old stuff below
// 	wa.ID := "2"
// 	wa.Brand := "Jaeger LeCoultre"

// 	// _, err = db.Exec("INSERT INTO watches (ID, Brand) VALUES ($1,$2)", info.ID, info.Brand)
// 	if err != nil {
// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 	return
// 	}

// 	fmt.Fprintf(w, "Record Created: ")
// 	fmt.Fprintf(w, "%s %sn", info.ID, info.Brand) // (3)
// 	}

//this is handling an error and can be called it in other page functions
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println("Index did not work. error in index")
	}
}

//This function takes id # as input and outputs watchid and corresponding brand

// func getWatch(id string) (watch watchInfo, err error) {
// db := dbConn()
// defer db.Close()
// if err != nil {
// 	log.Fatal(err)
// }
// err = db.Ping()
// if err != nil {
// 	fmt.Println(err.Error())
// }

// // do i need to add ";" at end of below function?
// // check for naming convention in mysql vs code vs frontend
// //sW (select watches)
// sW, err := db.Query("SELECT * FROM watches order by id")

// wa := watchInfo{}
// waS := []watchInfo{}

// for sW.Next() {
// 	var id int
// 	var brand string
// 	err = sW.Scan(&id, &brand)
// 	if err != nil {
// 		fmt.Println("sW.Scan didn't work")
// 	}
// 	wa.ID = id
// 	wa.Brand = brand
// 	waS = append(waS, wa)

// }
// tpl.ExecuteTemplate(w, "pic.gohtml", waS)

// err = db.QueryRow("SELECT id, brand from watches where id = ?;", id).Scan(&watch.ID, &watch.Brand)
// if err != nil {
// 	panic(err.Error())
// }

// return watch, err
// }
