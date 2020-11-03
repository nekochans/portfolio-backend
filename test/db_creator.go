package test

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/config"
)

type DbCreator struct{}

func (c *DbCreator) Create() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.GetTestDsn())

	if err != nil {
		return nil, err
	}

	return db, nil
}
