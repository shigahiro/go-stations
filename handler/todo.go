package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var reqBody model.CreateTODORequest
		if err := decoder.Decode(&reqBody); err != nil {
			return
		}

		if reqBody.Subject == "" {
			http.Error(w, "400", http.StatusBadRequest)
			return
		}

		todo, err := h.svc.CreateTODO(r.Context(), reqBody.Subject, reqBody.Description)

		if err != nil {
			return
		}

		encoder := json.NewEncoder(w)

		if err := encoder.Encode(model.CreateTODOResponse{TODO: *todo}); err != nil {
			return
		}
	}

	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var reqBody model.UpdateTODORequest
		if err := decoder.Decode(&reqBody); err != nil {
			return
		}

		if reqBody.ID == 0 || reqBody.Subject == "" {
			http.Error(w, "400", http.StatusBadRequest)
			return
		}

		todo, err := h.svc.UpdateTODO(r.Context(), reqBody.ID, reqBody.Subject, reqBody.Description)

		if err != nil {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		encoder := json.NewEncoder(w)

		if err := encoder.Encode(model.UpdateTODOResponse{TODO: *todo}); err != nil {
			return
		}
	}

}
