package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mirhijinam/wb-l0/internal/models"
)

type service interface {
	GetFromCache(uid string) (models.OrderData, error)
	PutIntoCache(ctx context.Context) error
}

type API struct {
	service service
}

func New(s service) API {
	return API{
		service: s,
	}
}

func (api API) GetOrderById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	data, err := api.service.GetFromCache(id)
	if err != nil {
		http.Error(w, "data not found", http.StatusNotFound)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}
