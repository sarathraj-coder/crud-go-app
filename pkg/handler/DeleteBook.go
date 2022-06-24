package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample_crud/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

	// 		break
	// 	}
	// }
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Error in deleting")
	} else {
		h.DB.Delete(&book)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
	}

}
