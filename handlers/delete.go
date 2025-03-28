package handlers

import (
	"net/http"
	"strconv"

	"github.com/RonaldAG/go-microservice/data"
	"github.com/gorilla/mux"
)

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE product")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Enable to convert the ID to int", http.StatusBadRequest)		
		return 
	}

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(rw, "Something wrong during deletion from db", http.StatusInternalServerError)
		return
	}
	
	rw.WriteHeader(http.StatusNoContent)
} 