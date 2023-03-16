package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	application "tempest-data-service-deta.space/pkg/application/entities"
)

func (sp *StorageProvider) GetAllFileInformation(ctx context.Context, username string) ([]application.File, error) {

	url := fmt.Sprintf("%s/%s/%s/query", sp.APIURLBase, sp.ProjectName, sp.DatabaseName)

	if username == "" {
		return []application.File{}, fmt.Errorf(`"username" is empty`)
	}

	body := bytes.Buffer{}
	err := json.NewEncoder(&body).Encode(
		newQueryAllItems(
			queryParams{
				User: username,
			},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error encoding body, err %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		return nil, fmt.Errorf("error building request, err %v", err)
	}

	request.Header.Add(headerAPIKey, sp.Key)
	request.Header.Add(headerContentType, valueApplicationJSON)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error calling service, err %v", err)
	}

	return readQueryReponse(*resp)
}

func (sp *StorageProvider) GetFileContent(ctx context.Context, username string, item string) (interface{}, error) {
	return nil, nil
}

func (sp *StorageProvider) UploadSmallFile(ctx context.Context, username string, itemName string, itemContent interface{}) error {
	return nil
}
