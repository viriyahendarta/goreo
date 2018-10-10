package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	cont "github.com/viriyahendarta/goreo/infra/context"
	e "github.com/viriyahendarta/goreo/infra/error"
	businessresource "github.com/viriyahendarta/goreo/resource/business"
)

//Resource holds data needed for service
type Resource struct {
	BusinessResource *businessresource.Resource
	Router           *mux.Router
}

//RenderJSON inject header and payload into response
func (r *Resource) RenderJSON(ctx context.Context, w http.ResponseWriter, data interface{}, err error) {
	var header APIHeader

	header.ProcessTime = cont.GetElapsedTime(ctx)
	status := http.StatusOK

	if err != nil {
		e := e.Cast(err)
		status = e.GetErrorTypeCode().HttpCode
		header.ErrorCode = e.GetErrorTypeCode().Message
		header.Message = e.GetMessage()
		header.Reason = e.Error()
		data = nil
	}

	response := APIResponse{
		Header:  header,
		Payload: data,
	}
	byteData, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(byteData)
}
