package test

import (
	"database/sql"
	"testing"

	"github.com/nekochans/portfolio-backend/config"
)

type DbCreator struct{}

func (c *DbCreator) Create(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", config.GetTestDsn())

	if err != nil {
		t.Fatal("Db Connect Error", err)
	}

	return db
}
