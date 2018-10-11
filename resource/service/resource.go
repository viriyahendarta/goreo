package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/viriyahendarta/goreo/config"

	"github.com/gorilla/mux"

	cont "github.com/viriyahendarta/goreo/infra/context"
	"github.com/viriyahendarta/goreo/infra/rror"
	businessresource "github.com/viriyahendarta/goreo/resource/business"
)

//Resource holds data needed for service
type Resource struct {
	BusinessResource *businessresource.Resource
	Router           *mux.Router
}

func (r *Resource) RenderJSON(ctx context.Context, w http.ResponseWriter, data interface{}, status int, err error) {
	header := APIHeader{
		ProcessTime: cont.GetElapsedTime(ctx),
		Message:     "success",
	}

	if err != nil {
		r := rror.Cast(err)
		header.Message = "error"

		if config.Get().Debug {
			header.Message = r.Message
			header.ErrorCode = int(r.Code)
			header.ErrorType = string(r.Type)
			header.ErrorReason = r.Err.Error()
		}
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
