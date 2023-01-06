package main

import(
	"encoding/json" // We are working with JSON so we need to bring this inbuilt package
	"log" //logs errors
	"net/http" //working with http -> used to create Api's
	"math/rand" // to generate random number
	"strconv" // will be used to used to convert integer to string
	"github.com/gorilla/mux"
)

// A struct is like a class in object oriented programming. 
// It has properties and methods like one in Java and C++


// Book Struct (Model)
type Book struct {
	ID  string  `json:"id"`
	Isbn  string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author  `json:"author"`

}

// Author Struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}


//Init books var as a slice Book struct
//a slice is an array with variable length
var books []Book


//every function that we create that is a rout handler has to have
//these two parameters
//get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get a single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//getting the book by searching id
	params := mux.Vars(r) //get parameters
	// Loop through books and find with ID
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	} 
	json.NewEncoder(w).Encode(&Book{})
}

//create a new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_= json.NewDecoder(r.Body).Decode(&book) 
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}

//update a book
func updateBook(w http.ResponseWriter, r *http.Request){

}

//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){

}





func main(){
	//Init the mux router 
	r := mux.NewRouter()

	//Mock Data 
	books = append(books, Book{ID: "1", Isbn: "4773", Title: "A Freshers Guide to Devtron", Author: &Author{Firstname: "Shashwat", Lastname: "Dadhich"	} })
	books = append(books, Book{ID: "2", Isbn: "4273", Title: "Dockers", Author: &Author{Firstname: "Steve", Lastname: "Smith"} })


	//creating router handlers which will establish endpoints for our api's
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))//To run the server

}