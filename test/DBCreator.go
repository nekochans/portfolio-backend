package test

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/config"
	"testing"
)

type DBCreator struct{}

func (d *DBCreator) Create(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", config.GetTestDsn())

	if err != nil {
		t.Fatal("DB Connect Error", err)
	}

	return db
}
