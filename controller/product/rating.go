package product

import (
	"encoding/json"
	"fmt"
	"go-learn/entities"
	"go-learn/library/response"
	"go-learn/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type _ControllerProductRating struct {
	service service.Service
}

func NewControllerProductRating(service service.Service) *_ControllerProductRating {
	return &_ControllerProductRating{service: service}
}

func (c *_ControllerProductRating) Rating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		payload     entities.PayloadRating
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success Rating")
		bearer = r.Header.Get("Authorization")
	)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	if err := payload.Validate(); err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	rawID := mux.Vars(r)["id"]
	if rawID == ":id" {
		response := *errResponse.WithError("ID cannot be empty!")
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(output)
		return
	}
	id, _ := uuid.Parse(rawID)
	if err := c.service.ProductService.Rating(id, &payload, bearer); err != nil {
		fmt.Println("errr", err)
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
	w.WriteHeader(http.StatusCreated)
	w.Write(object)
}
