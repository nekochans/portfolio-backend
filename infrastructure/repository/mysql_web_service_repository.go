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

func (r *MysqlWebServiceRepository) FindAll() (domain.WebServices, error) {
	query := `
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

	stmt, ErrPrepare := r.Db.Prepare(query)
	if ErrPrepare != nil {
		ErrBackend := &domain.BackendError{Message: "Db.Prepare Error", Err: ErrPrepare}
		return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", ErrBackend)
	}

	defer func() {
		ErrStmtClose := stmt.Close()
		if ErrStmtClose != nil {
			log.Fatal(ErrStmtClose, "stmt.Close() Fatal.")
		}
	}()

	rows, ErrQuery := stmt.Query()

	if ErrQuery != nil {
		ErrBackend := &domain.BackendError{Message: "stmt.Query Error", Err: ErrQuery}
		return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", ErrBackend)
	}

	var tableData WebServiceFindAllTableData
	var webServices domain.WebServices
	for rows.Next() {
		ErrRowsScan := rows.Scan(&tableData.Id, &tableData.Url, &tableData.Description)
		webServices = append(
			webServices,
			&Openapi.WebService{
				Id:          tableData.Id,
				Url:         tableData.Url,
				Description: tableData.Description,
			},
		)

		if ErrRowsScan != nil {
			ErrBackend := &domain.BackendError{Message: "rows.Scan Error", Err: ErrRowsScan}
			return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", ErrBackend)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		ErrBackend := &domain.BackendError{Message: "WebServices Not Found"}
		return nil, xerrors.Errorf("MysqlWebServiceRepository.FindAll: %w", ErrBackend)
	}

	return webServices, nil
}
