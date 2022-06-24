package main

import (
	"fmt"
	"net/http"
	"sample_crud/pkg/db"
	"sample_crud/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {

	DB := db.Init()
	h := handler.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/", test_handler)
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/book", h.AddBookHandler).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", h.UpdateBookHandler).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", h.DeleteBook).Methods(http.MethodDelete)
	fmt.Println("Starting server")
	http.ListenAndServe(":8000", router)
}

func test_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is testing handler")
}
