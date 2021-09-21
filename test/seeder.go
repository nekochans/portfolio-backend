package test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Seeder struct {
	Db      *sql.DB
	DirPath string
}

func (s *Seeder) Execute() error {
	files, ErrReadDir := ioutil.ReadDir(s.DirPath)
	if ErrReadDir != nil {
		return errors.Wrap(ErrReadDir, "failed to read dir")
	}

	tx, ErrTransactionBegin := s.Db.Begin()
	if ErrTransactionBegin != nil {
		return errors.Wrap(ErrTransactionBegin, "failed to begin transaction")
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ".csv" {
			continue
		}

		table := file.Name()[:len(file.Name())-len(ext)]
		csvFilePath := filepath.Join(s.DirPath, file.Name())

		if _, ErrLoadData := s.loadDataFromCsv(tx, table, csvFilePath); ErrLoadData != nil {
			ErrRollback := tx.Rollback()
			if ErrRollback != nil {
				return errors.Wrap(ErrRollback, "failed to Transaction.Rollback()")
			}
			return errors.Wrap(ErrLoadData, "failed to loadDataFromCsv")
		}
	}

	return tx.Commit()
}

func (s *Seeder) TruncateAllTable() error {
	tx, ErrTransactionBegin := s.Db.Begin()
	if ErrTransactionBegin != nil {
		return errors.Wrap(ErrTransactionBegin, "failed to s.Db.Begin()")
	}

	_, ErrSetForeignKeyFalse := tx.Exec("SET FOREIGN_KEY_CHECKS=0")
	if ErrSetForeignKeyFalse != nil {
		ErrRollback := tx.Rollback()
		if ErrRollback != nil {
			return errors.Wrap(ErrRollback, "failed to Transaction.Rollback()")
		}
		return errors.Wrap(ErrSetForeignKeyFalse, "failed to exec sql")
	}

	// TODO テーブル分ループさせるように改修を行う
	_, ErrTruncateMembers := tx.Exec("TRUNCATE members")
	if ErrTruncateMembers != nil {
		ErrRollback := tx.Rollback()
		if ErrRollback != nil {
			return errors.Wrap(ErrRollback, "failed to Transaction.Rollback()")
		}
		return errors.Wrap(ErrTruncateMembers, "failed to exec sql")
	}

	_, ErrTruncateGitHubUsers := tx.Exec("TRUNCATE members_github_users")
	if ErrTruncateGitHubUsers != nil {
		ErrRollback := tx.Rollback()
		if ErrRollback != nil {
			return errors.Wrap(ErrRollback, "failed to Transaction.Rollback()")
		}
		return errors.Wrap(ErrTruncateGitHubUsers, "failed to exec sql")
	}

	_, ErrTruncateWebServices := tx.Exec("TRUNCATE webservices")
	if ErrTruncateWebServices != nil {
		ErrRollback := tx.Rollback()
		if ErrRollback != nil {
			return errors.Wrap(ErrRollback, "failed to Transaction.Rollback()")
		}
		return errors.Wrap(ErrTruncateWebServices, "failed to exec sql")
	}

	_, ErrSetForeignKeyTrue := tx.Exec("SET FOREIGN_KEY_CHECKS=1")
	if ErrSetForeignKeyTrue != nil {
		ErrRollback := tx.Rollback()
		if ErrRollback != nil {
			return errors.Wrap(ErrRollback, "failed to Transaction.Rollback()")
		}
		return errors.Wrap(ErrSetForeignKeyTrue, "failed to exec sql")
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
