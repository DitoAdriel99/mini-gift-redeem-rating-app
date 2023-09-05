package product

import (
	"encoding/json"
	"go-learn/entities"
	"go-learn/library/response"
	"go-learn/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type _ControllerProductUpdate struct {
	service service.Service
}

func NewControllerProductUpdate(service service.Service) *_ControllerProductUpdate {
	return &_ControllerProductUpdate{service: service}
}

func (c *_ControllerProductUpdate) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
		payload entities.Product
	)
	rawID := mux.Vars(r)["id"]
	if rawID == ":id" {
		response := *errResponse.WithError("ID cannot be empty!")
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(output)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	id, _ := uuid.Parse(rawID)
	if err := c.service.ProductService.Update(id, &payload); err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	object, err := json.Marshal(succResponse)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(object)
}
