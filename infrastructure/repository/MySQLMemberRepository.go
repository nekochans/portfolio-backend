package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"golang.org/x/xerrors"
)

type MySQLMemberRepository struct {
	DB *sql.DB
}

type FindTableData struct {
	Id   int64
	GithubId   string
	AvatarUrl  string
	CvRepoName string
}

func (m *MySQLMemberRepository) Find(id int) (*Openapi.Member, error) {
	sql := `
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

	stmt, err := m.DB.Prepare(sql)

	if err != nil {
		appErr := &domain.BackendError{Msg: "DB.Prepare Error", Err: err}
		return nil, xerrors.Errorf("MySQLMemberRepository.Find: %w", appErr)
	}

	defer stmt.Close()

	var tableData FindTableData

	err = stmt.QueryRow(id).Scan(&tableData.Id, &tableData.GithubId, &tableData.AvatarUrl, &tableData.CvRepoName)
	if err != nil {
		// この条件の時はデータが1件も存在しない
		if err.Error() == "sql: no rows in result set" {
			appErr := &domain.BackendError{Msg: "Member Not Found"}
			return nil, xerrors.Errorf("MySQLMemberRepository.Find: %w", appErr)
		}

		appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: err}
		return nil, xerrors.Errorf("MySQLMemberRepository.Find: %w", appErr)
	}

	me := &Openapi.Member{
		Id:             tableData.Id,
		GithubUserName: tableData.GithubId,
		GithubPicture:  tableData.AvatarUrl,
		CvUrl:          "https://github.com/" + tableData.GithubId + "/" + tableData.CvRepoName,
	}

	return me, nil
}

type FindAllTableData struct {
	Id   int64
	GithubId   string
	AvatarUrl  string
	CvRepoName string
}

func (m *MySQLMemberRepository) FindAll() (domain.Members, error) {
	sql := `
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

	stmt, err := m.DB.Prepare(sql)
	if err != nil {
		appErr := &domain.BackendError{Msg: "DB.Prepare Error", Err: err}
		return nil, xerrors.Errorf("MySQLMemberRepository.FindAll: %w", appErr)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		appErr := &domain.BackendError{Msg: "stmt.Query Error", Err: err}
		return nil, xerrors.Errorf("MySQLMemberRepository.FindAll: %w", appErr)
	}

	var tableData FindAllTableData
	var ms domain.Members
	for rows.Next() {
		err := rows.Scan(&tableData.Id, &tableData.GithubId, &tableData.AvatarUrl, &tableData.CvRepoName)
		ms = append(
			ms,
			&Openapi.Member{
				Id:             tableData.Id,
				GithubUserName: tableData.GithubId,
				GithubPicture:  tableData.AvatarUrl,
				CvUrl:          "https://github.com/" + tableData.GithubId + "/" + tableData.CvRepoName,
			},
		)

		if err != nil {
			appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: err}
			return nil, xerrors.Errorf("MySQLMemberRepository.FindAll: %w", appErr)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		appErr := &domain.BackendError{Msg: "Members Not Found"}
		return nil, xerrors.Errorf("MySQLMemberRepository.FindAll: %w", appErr)
	}

	return ms, nil
}
