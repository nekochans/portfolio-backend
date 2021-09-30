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

	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return nil, errors.Wrap(domain.ErrWebServiceRepositoryUnexpected, err.Error())
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			r.Logger.Error(
				"stmt.Close() Fatal.",
				zap.Error(
					errors.Wrap(err, "stmt.Close() Fatal."),
				),
			)
		}
	}()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(domain.ErrWebServiceRepositoryUnexpected, err.Error())
	}

	var tableData WebServiceFindAllTableData
	var webServices domain.WebServices
	for rows.Next() {
		if err := rows.Scan(&tableData.Id, &tableData.Url, &tableData.Description); err != nil {
			return nil, errors.Wrap(domain.ErrWebServiceRepositoryUnexpected, err.Error())
		}

		webServices = append(
			webServices,
			&Openapi.WebService{
				Id:          tableData.Id,
				Url:         tableData.Url,
				Description: tableData.Description,
			},
		)
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		return nil, errors.Wrap(domain.ErrWebServiceNotFound, "record not found in webservices")
	}

	return webServices, nil
}
