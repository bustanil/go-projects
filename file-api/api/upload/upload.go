package upload

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"bustanil.com/file-api/dto"
	"bustanil.com/file-api/handler/upload"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3client s3.PresignClient
)

type API struct {
	handler upload.Handler
}

func NewAPI(handler upload.Handler) *API {
	return &API{handler: handler}
}

func (a *API) PostFileMetadata(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	mdReq := &dto.PostFileMetadataRequest{}
	err = json.Unmarshal(body, &mdReq)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mdResp, err := a.handler.HandleUpload(r.Context(), mdReq)
	if err != nil {
		log.Printf("failed to handle upload %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(mdResp)
	if err != nil {
		log.Printf("Error serializing request: %+v", err)
		return
	}

	w.Write(json)
}
