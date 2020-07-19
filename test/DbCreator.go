package test

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/config"
	"testing"
)

type DbCreator struct{}

func (d *DbCreator) Create(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", config.GetTestDsn())

	if err != nil {
		t.Fatal("Db Connect Error", err)
	}

	return db
}
