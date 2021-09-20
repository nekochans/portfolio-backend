package repository

import (
	"database/sql"

	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MysqlWebServiceRepository struct {
	Db     *sql.DB
	Logger *zap.Logger
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
		return nil, errors.Wrap(domain.ErrWebServiceRepositoryUnexpected, ErrPrepare.Error())
	}

	defer func() {
		ErrStmtClose := stmt.Close()
		if ErrStmtClose != nil {
			r.Logger.Error("stmt.Close() Fatal.", zap.Error(ErrStmtClose))
		}
	}()

	rows, ErrQuery := stmt.Query()

	if ErrQuery != nil {
		return nil, errors.Wrap(domain.ErrWebServiceRepositoryUnexpected, ErrQuery.Error())
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
			return nil, errors.Wrap(domain.ErrWebServiceRepositoryUnexpected, ErrRowsScan.Error())
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		return nil, errors.Wrap(domain.ErrWebServiceNotFound, "record not found in webservices")
	}

	return webServices, nil
}
