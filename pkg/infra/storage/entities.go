package storage

import (
	"fmt"

	application "tempest-data-service-deta.space/pkg/application/entities"
)

type queryBodyReponse struct {
	Items []application.File `json:"items"`
}

func storageResponseToApplicationResponse(queryAllItems queryBodyReponse) ([]application.File, error) {

	if len(queryAllItems.Items) == 0 {
		return []application.File{}, fmt.Errorf("could not convert response, empty")
	}

	resp := []application.File{}
	for _, val := range queryAllItems.Items {
		resp = append(resp,
			application.File{
				Key:      val.Key,
				User:     val.User,
				Metadata: val.Metadata,
				Data:     val.Data,
			},
		)
	}

	return resp, nil
}

type postNewItem struct {
	Item application.File `json:"item"`
}

type queryParams struct {
	User string `json:"username"`
}

type queryAllItems struct {
	Query []queryParams `json:"query"`
}

func newQueryAllItems(param queryParams) queryAllItems {
	return queryAllItems{
		Query: []queryParams{
			param,
		},
	}
}
