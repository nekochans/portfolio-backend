package test

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"path/filepath"
)

type Seeder struct {
	DB      *sql.DB
	DirPath string
}

func (s *Seeder) Execute() error {
	files, err := ioutil.ReadDir(s.DirPath)
	if err != nil {
		return err
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ".csv" {
			continue
		}

		table := file.Name()[:len(file.Name())-len(ext)]
		csvFilePath := filepath.Join(s.DirPath, file.Name())

		if _, err := loadDataFromCSV(tx, table, csvFilePath); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *Seeder) TruncateAllTable() error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("SET FOREIGN_KEY_CHECKS=0")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("TRUNCATE members")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("TRUNCATE members_github_users")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("SET FOREIGN_KEY_CHECKS=1")
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func loadDataFromCSV(tx *sql.Tx, table, filePath string) (sql.Result, error) {
	s := `
		LOAD DATA
			LOCAL INFILE '%s'
		INTO TABLE %s
		FIELDS
			TERMINATED BY ','
		LINES
			TERMINATED BY '\n'
			IGNORE 1 LINES
	`

	mysql.RegisterLocalFile(filePath)

	return tx.Exec(fmt.Sprintf(s, filePath, table))
}
