package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	application "tempest-data-service-deta.space/pkg/application/entities"
)

const (
	headerAPIKey         = "X-API-Key"
	headerContentType    = "Content-Type"
	valueApplicationJSON = "application/json"
)

func readQueryReponse(resp http.Response) ([]application.File, error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []application.File{}, fmt.Errorf("error reading body, err %v", err)
	}

	applicationResponse := queryBodyReponse{}
	err = json.Unmarshal(body, &applicationResponse)
	if err != nil {
		return []application.File{}, fmt.Errorf("error unmarshalling body, err %v", err)
	}

	return storageResponseToApplicationResponse(applicationResponse)
}
