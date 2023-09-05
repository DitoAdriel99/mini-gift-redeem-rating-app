package product

import (
	"encoding/json"
	"go-learn/library/meta"
	"go-learn/library/response"
	"go-learn/service"
	"net/http"
)

type _ControllerProductGetAll struct {
	service service.Service
}

func NewControllerProductGetAll(service service.Service) *_ControllerProductGetAll {
	return &_ControllerProductGetAll{service: service}
}

func (c *_ControllerProductGetAll) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
		query    = r.URL.Query()
		metadata = meta.MetadataFromURL(query)
	)

	data, err := c.service.ProductService.GetAll(&metadata)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	metadata.Total = int64(len(data))
	response := *succResponse.WithData(data).WithMeta(metadata)
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
