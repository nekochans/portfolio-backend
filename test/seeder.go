package test

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Seeder struct {
	Db      *sql.DB
	DirPath string
}

var (
	ErrReadTestDataDir     = errors.New("failed to read test data dir")
	ErrBeginTransaction    = errors.New("failed to begin transaction")
	ErrTransactionRollback = errors.New("failed to transaction rollback")
	ErrLoadDataFromCsv     = errors.New("failed to load data from csv")
	ErrExecSql             = errors.New("failed to exec sql")
)

func (s *Seeder) Execute() error {
	files, err := os.ReadDir(s.DirPath)
	if err != nil {
		return errors.Wrap(ErrReadTestDataDir, err.Error())
	}

	tx, err := s.Db.Begin()
	if err != nil {
		return errors.Wrap(ErrBeginTransaction, err.Error())
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ".csv" {
			continue
		}

		table := file.Name()[:len(file.Name())-len(ext)]
		csvFilePath := filepath.Join(s.DirPath, file.Name())

		if _, err := s.loadDataFromCsv(tx, table, csvFilePath); err != nil {
			if err := tx.Rollback(); err != nil {
				return errors.Wrap(ErrTransactionRollback, err.Error())
			}

			return errors.Wrap(ErrLoadDataFromCsv, err.Error())
		}
	}

	return tx.Commit()
}

func (s *Seeder) TruncateAllTable() error {
	tx, err := s.Db.Begin()
	if err != nil {
		return errors.Wrap(ErrBeginTransaction, err.Error())
	}

	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS=0"); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(ErrTransactionRollback, err.Error())
		}

		return errors.Wrap(ErrExecSql, err.Error())
	}

	// TODO テーブル分ループさせるように改修を行う
	if _, err := tx.Exec("TRUNCATE members"); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(ErrTransactionRollback, err.Error())
		}

		return errors.Wrap(ErrExecSql, err.Error())
	}

	if _, err := tx.Exec("TRUNCATE members_github_users"); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(ErrTransactionRollback, err.Error())
		}

		return errors.Wrap(ErrExecSql, err.Error())
	}

	if _, err := tx.Exec("TRUNCATE webservices"); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(ErrTransactionRollback, err.Error())
		}

		return errors.Wrap(ErrExecSql, err.Error())
	}

	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS=1"); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(ErrTransactionRollback, err.Error())
		}

		return errors.Wrap(ErrExecSql, err.Error())
	}

	return tx.Commit()
}

func (s *Seeder) loadDataFromCsv(tx *sql.Tx, table, filePath string) (sql.Result, error) {
	query := `
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

	return tx.Exec(fmt.Sprintf(query, filePath, table))
}
