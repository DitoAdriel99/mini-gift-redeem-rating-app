package product

import (
	"encoding/json"
	"go-learn/library/response"
	"go-learn/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type _ControllerProductDetail struct {
	service service.Service
}

func NewControllerProductDetail(service service.Service) *_ControllerProductDetail {
	return &_ControllerProductDetail{service: service}
}

func (c *_ControllerProductDetail) Detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
	)
	rawID := mux.Vars(r)["id"]
	if rawID == ":id" {
		response := *errResponse.WithError("ID cannot be empty!")
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(output)
		return
	}
	id, _ := uuid.Parse(rawID)
	data, err := c.service.ProductService.Detail(id)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	response := *succResponse.WithData(data)
	object, err := json.Marshal(response)
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
