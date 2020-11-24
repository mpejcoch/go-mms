package server

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMockService() (*Service, sqlmock.Sqlmock, error) {
	cacheDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create mock cache DB: %s", err)
	}

	templates := CreateTemplates()
	webService := NewService(templates, cacheDB)

	return webService, mock, nil
}
