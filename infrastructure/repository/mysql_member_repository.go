package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MysqlMemberRepository struct {
	Db     *sql.DB
	Logger *zap.Logger
}

type FindTableData struct {
	Id         int64
	GithubId   string
	AvatarUrl  string
	CvRepoName string
}

func (r *MysqlMemberRepository) Find(id int) (*Openapi.Member, error) {
	query := `
		SELECT
		  m.id AS Id,
		  mgu.github_id AS GithubId,
		  mgu.avatar_url AS AvatarUrl,
		  mgu.cv_repo_name AS CvRepoName
		FROM
		  members AS m
		INNER JOIN
		  members_github_users AS mgu
		ON
		  m.id = mgu.member_id
		WHERE
			m.id = ?
	`

	stmt, ErrPrepare := r.Db.Prepare(query)

	if ErrPrepare != nil {
		return nil, errors.Wrap(domain.ErrMemberRepositoryUnexpected, ErrPrepare.Error())
	}

	defer func() {
		ErrStmtClose := stmt.Close()
		if ErrStmtClose != nil {
			r.Logger.Error("stmt.Close() Fatal.", zap.Error(ErrStmtClose))
		}
	}()

	var tableData FindTableData

	ErrQuery := stmt.QueryRow(id).Scan(&tableData.Id, &tableData.GithubId, &tableData.AvatarUrl, &tableData.CvRepoName)
	if ErrQuery != nil {
		// この条件の時はデータが1件も存在しない
		if ErrQuery.Error() == "sql: no rows in result set" {
			return nil, errors.Wrap(domain.ErrMemberNotFound, ErrQuery.Error())
		}

		return nil, errors.Wrap(domain.ErrMemberRepositoryUnexpected, ErrQuery.Error())
	}

	member := &Openapi.Member{
		Id:             tableData.Id,
		GithubUserName: tableData.GithubId,
		GithubPicture:  tableData.AvatarUrl,
		CvUrl:          "https://github.com/" + tableData.GithubId + "/" + tableData.CvRepoName,
	}

	return member, nil
}

type FindAllTableData struct {
	Id         int64
	GithubId   string
	AvatarUrl  string
	CvRepoName string
}

func (r *MysqlMemberRepository) FindAll() (domain.Members, error) {
	query := `
		SELECT
		  m.id AS Id,
		  mgu.github_id AS GithubId,
		  mgu.avatar_url AS AvatarUrl,
		  mgu.cv_repo_name AS CvRepoName
		FROM
		  members AS m
		INNER JOIN
		  members_github_users AS mgu
		ON
		  m.id = mgu.member_id
		ORDER BY
			m.id
		ASC
	`

	stmt, ErrPrepare := r.Db.Prepare(query)
	if ErrPrepare != nil {
		return nil, errors.Wrap(domain.ErrMemberRepositoryUnexpected, ErrPrepare.Error())
	}

	defer func() {
		ErrStmtClose := stmt.Close()
		if ErrStmtClose != nil {
			r.Logger.Error("stmt.Close() Fatal.", zap.Error(ErrStmtClose))
		}
	}()

	rows, ErrQuery := stmt.Query()

	if ErrQuery != nil {
		return nil, errors.Wrap(domain.ErrMemberRepositoryUnexpected, ErrQuery.Error())
	}

	var tableData FindAllTableData
	var members domain.Members
	for rows.Next() {
		ErrRowsScan := rows.Scan(&tableData.Id, &tableData.GithubId, &tableData.AvatarUrl, &tableData.CvRepoName)
		members = append(
			members,
			&Openapi.Member{
				Id:             tableData.Id,
				GithubUserName: tableData.GithubId,
				GithubPicture:  tableData.AvatarUrl,
				CvUrl:          "https://github.com/" + tableData.GithubId + "/" + tableData.CvRepoName,
			},
		)

		if ErrRowsScan != nil {
			return nil, errors.Wrap(domain.ErrMemberRepositoryUnexpected, ErrRowsScan.Error())
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		return nil, errors.Wrap(domain.ErrMemberNotFound, "record not found in members")
	}

	return members, nil
}
