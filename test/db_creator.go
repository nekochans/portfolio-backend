package test

import (
	"database/sql"

	"github.com/nekochans/portfolio-backend/config"
	"github.com/pkg/errors"
)

type DbCreator struct{}

var ErrTestDbConnection = errors.New("failed to connect to test mysql server")

func (c *DbCreator) Create() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.GetTestDsn())

	if err != nil {
		return nil, errors.Wrap(ErrTestDbConnection, err.Error())
	}

	return db, nil
}
