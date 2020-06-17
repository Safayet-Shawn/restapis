package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MyURL represents the model for Link shorten api
type MyURL struct {
	ID       string `json:"id"`
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.Exec("CREATE DATABASE ShortUrl")
	db.Exec("USE ShortUrl")
	db.AutoMigrate(&MyURL{})
}

func getUrls(w http.ResponseWriter, req *http.Request) {

}

func createURL(w http.ResponseWriter, r *http.Request) {

}

func getURL(w http.ResponseWriter, req *http.Request) {

}

func deleteURL(w http.ResponseWriter, req *http.Request) {

}
func main() {
	router := mux.NewRouter()
	// Read all
	router.HandleFunc("/myurls", getUrls).Methods("POST")
	// Create
	router.HandleFunc("/myurl/{orderId}", getURL).Methods("GET")
	// create url
	router.HandleFunc("myurl/create", createURL).Methods("POST")

	// Delete url
	router.HandleFunc("/myurl/{orderId}", deleteURL).Methods("DELETE")
	//  db connection
	initDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
