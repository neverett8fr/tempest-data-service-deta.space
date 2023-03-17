package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

func (sp *StorageProvider) GetFileContent(ctx context.Context, username string, key string) (application.File, error) {

	url := fmt.Sprintf("%s/%s/%s/query", sp.APIURLBase, sp.ProjectName, sp.DatabaseName)

	if username == "" {
		return application.File{}, fmt.Errorf(`"username" is empty`)
	}

	body := bytes.Buffer{}
	err := json.NewEncoder(&body).Encode(
		newQuerySpecificItem(
			queryItemParam{
				User: username,
				Key:  key,
			},
		),
	)
	if err != nil {
		return application.File{}, fmt.Errorf("error encoding body, err %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		return application.File{}, fmt.Errorf("error building request, err %v", err)
	}

	request.Header.Add(headerAPIKey, sp.Key)
	request.Header.Add(headerContentType, valueApplicationJSON)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return application.File{}, fmt.Errorf("error calling service, err %v", err)
	}

	query, err := readQueryReponse(*resp)
	return query[0], err

}

func (sp *StorageProvider) UploadSmallFile(ctx context.Context, username string, fileName string, fileExt string, fileSize int, fileContent []byte) error {

	url := fmt.Sprintf("%s/%s/%s/items", sp.APIURLBase, sp.ProjectName, sp.DatabaseName)

	if username == "" {
		return fmt.Errorf(`"username" is empty`)
	}

	body := bytes.Buffer{}
	err := json.NewEncoder(&body).Encode(
		postNewItem{
			Item: application.File{
				Key:  uuid.NewString(),
				User: username,
				Metadata: application.FileMetadata{
					Extension: fileExt,
					Name:      fileName,
					Size:      fileSize,
				},
				Data: fileContent,
			},
		},
	)

	if err != nil {
		return fmt.Errorf("error encoding body, err %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		return fmt.Errorf("error building request, err %v", err)
	}

	request.Header.Add(headerAPIKey, sp.Key)
	request.Header.Add(headerContentType, valueApplicationJSON)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error calling service, err %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error calling service, status received %v", resp.StatusCode)
	}

	return nil

}
