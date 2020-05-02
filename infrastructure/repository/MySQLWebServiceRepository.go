package repository

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/domain"
	"golang.org/x/xerrors"
)

type MySQLWebServiceRepository struct {
	DB *sql.DB
}

type WebServiceFindAllTableData struct {
	ID          int
	URL         string
	Description string
}

func (m *MySQLWebServiceRepository) FindAll() (domain.WebServices, error) {
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

	stmt, err := m.DB.Prepare(sql)
	if err != nil {
		appErr := &domain.BackendError{Msg: "DB.Prepare Error", Err: err}
		return nil, xerrors.Errorf("MySQLWebServiceRepository.FindAll: %w", appErr)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		appErr := &domain.BackendError{Msg: "stmt.Query Error", Err: err}
		return nil, xerrors.Errorf("MySQLWebServiceRepository.FindAll: %w", appErr)
	}

	var tableData WebServiceFindAllTableData
	var ws domain.WebServices
	for rows.Next() {
		err := rows.Scan(&tableData.ID, &tableData.URL, &tableData.Description)
		ws = append(
			ws,
			&domain.WebService{
				ID:          tableData.ID,
				URL:         tableData.URL,
				Description: tableData.Description,
			},
		)

		if err != nil {
			appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: err}
			return nil, xerrors.Errorf("MySQLWebServiceRepository.FindAll: %w", appErr)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.ID == 0 {
		appErr := &domain.BackendError{Msg: "WebServices Not Found"}
		return nil, xerrors.Errorf("MySQLWebServiceRepository.FindAll: %w", appErr)
	}

	return ws, nil
}
