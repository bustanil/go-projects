package upload

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"bustanil.com/file-api/db"
	"bustanil.com/file-api/dto"
	"bustanil.com/file-api/handler/upload"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3client s3.PresignClient
)

type API struct {
	handler upload.Handler
}

func NewAPI(cfg *aws.Config, pg *db.Postgres) *API {
	return &API{
		handler: upload.NewHandler(cfg, pg),
	}
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
	}

	mdResp, err := a.handler.HandleUpload(r.Context(), mdReq)
	if err != nil {
		log.Printf("failed to handle upload %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	json, err := json.Marshal(mdResp)
	if err != nil {
		log.Printf("Error serializing request: %+v", err)
	}

	w.Write(json)
}
