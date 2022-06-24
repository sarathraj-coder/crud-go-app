package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sample_crud/pkg/models"
)

func (h handler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// book.Id = rand.Intn(100)
	// mocks.Books = append(mocks.Books, book)
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "applicatio/json")
		json.NewEncoder(w).Encode("Error in saving data")
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "applicatio/json")
		json.NewEncoder(w).Encode("Created")
	}

}
