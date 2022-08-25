package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"shapes/libs/util"
	"shapes/service"
)

type AreaService interface {
	InsertArea(ctx context.Context, req *service.AreaRequest) error
}

type AreaHTTPHandler struct {
	Service AreaService
}

func (a *AreaHTTPHandler) Insert(rw http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		body = service.AreaRequest{}
	)

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.HTTPResponse(rw, http.StatusBadRequest, "parse request body, an error occured", nil)
		return
	}

	if err := a.Service.InsertArea(ctx, &body); err != nil {
		util.ErrHTTPResponse(ctx, rw, err)
		return
	}

	util.HTTPResponse(rw, http.StatusOK, "success insert area", nil)
}
