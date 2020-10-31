package repository

import (
	"database/sql"
	"log"

	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"golang.org/x/xerrors"
)

type MysqlWebServiceRepository struct {
	Db *sql.DB
}

type WebServiceFindAllTableData struct {
	Id          int64
	Url         string
	Description string
}

func (m *MysqlWebServiceRepository) FindAll() (domain.WebServices, error) {
	sql := `
		SELECT
		  id,
		  url,
		  description
		FROM
		  webservices
		ORDER BY
			id
		ASC
	`

	stmt, ErrPrepare := m.Db.Prepare(sql)
	if ErrPrepare != nil {
		appErr := &domain.BackendError{Msg: "Db.Prepare Error", Err: ErrPrepare}
		return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", appErr)
	}

	defer func() {
		ErrStmtClose := stmt.Close()
		if ErrStmtClose != nil {
			log.Fatal(ErrStmtClose, "stmt.Close() Fatal.")
		}
	}()

	rows, ErrQuery := stmt.Query()

	if ErrQuery != nil {
		appErr := &domain.BackendError{Msg: "stmt.Query Error", Err: ErrQuery}
		return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", appErr)
	}

	var tableData WebServiceFindAllTableData
	var ws domain.WebServices
	for rows.Next() {
		ErrRowsScan := rows.Scan(&tableData.Id, &tableData.Url, &tableData.Description)
		ws = append(
			ws,
			&Openapi.WebService{
				Id:          tableData.Id,
				Url:         tableData.Url,
				Description: tableData.Description,
			},
		)

		if ErrRowsScan != nil {
			appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: ErrRowsScan}
			return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", appErr)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		appErr := &domain.BackendError{Msg: "WebServices Not Found"}
		return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", appErr)
	}

	return ws, nil
}
