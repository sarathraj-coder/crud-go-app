package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sample_crud/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Error in reading data")
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("No book found ")
	} else {
		book.Author = updatedBook.Author
		book.Desc = updatedBook.Desc
		book.Title = updatedBook.Title
		h.DB.Save(&book)

		//mocks.Books[index] = book
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Updated")
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(&book)
	}

	// for index, book := range mocks.Books {
	// 	if book.Id == id {

	// 		break
	// 	}
	// }

}
