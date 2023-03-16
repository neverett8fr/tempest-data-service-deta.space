package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	st "tempest-data-service-deta.space/pkg/infra/storage"

	"tempest-data-service-deta.space/pkg/config"

	"github.com/gorilla/mux"
)

var (
	StorageProvider st.StorageProvider
)

const (
	username = "username"
	item     = "item"
)

func NewServiceRoutes(r *mux.Router, conf config.Config) {
	sp, err := st.InitialiseStorageProvider(
		context.Background(),
		conf.Storage.APIURLBase,
		conf.Storage.ProjectName,
		conf.Storage.DatabaseName,
		conf.Storage.DatabaseKey,
	)
	if err != nil {
		log.Printf("error initialising storage provider, err %v", err)
	}

	StorageProvider = sp

	newDataInformation(r)
	newDataOperation(r)
}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
