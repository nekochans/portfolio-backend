package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"golang.org/x/xerrors"
)

type MysqlMemberRepository struct {
	Db *sql.DB
}

type FindTableData struct {
	Id         int64
	GithubId   string
	AvatarUrl  string
	CvRepoName string
}

func (m *MysqlMemberRepository) Find(id int) (*Openapi.Member, error) {
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

	stmt, ErrPrepare := m.Db.Prepare(sql)

	if ErrPrepare != nil {
		appErr := &domain.BackendError{Msg: "Db.Prepare Error", Err: ErrPrepare}
		return nil, xerrors.Errorf("MysqlMemberRepository.Find: %w", appErr)
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err, "stmt.Close() Fatal.")
		}
	}()

	var tableData FindTableData

	ErrQuery := stmt.QueryRow(id).Scan(&tableData.Id, &tableData.GithubId, &tableData.AvatarUrl, &tableData.CvRepoName)
	if ErrQuery != nil {
		// この条件の時はデータが1件も存在しない
		if ErrQuery.Error() == "sql: no rows in result set" {
			appErr := &domain.BackendError{Msg: "Member Not Found"}
			return nil, xerrors.Errorf("MysqlMemberRepository.Find: %w", appErr)
		}

		appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: ErrQuery}
		return nil, xerrors.Errorf("MysqlMemberRepository.Find: %w", appErr)
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
	Id         int64
	GithubId   string
	AvatarUrl  string
	CvRepoName string
}

func (m *MysqlMemberRepository) FindAll() (domain.Members, error) {
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

	stmt, ErrPrepare := m.Db.Prepare(sql)
	if ErrPrepare != nil {
		appErr := &domain.BackendError{Msg: "Db.Prepare Error", Err: ErrPrepare}
		return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", appErr)
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
		return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", appErr)
	}

	var tableData FindAllTableData
	var ms domain.Members
	for rows.Next() {
		ErrRowsScan := rows.Scan(&tableData.Id, &tableData.GithubId, &tableData.AvatarUrl, &tableData.CvRepoName)
		ms = append(
			ms,
			&Openapi.Member{
				Id:             tableData.Id,
				GithubUserName: tableData.GithubId,
				GithubPicture:  tableData.AvatarUrl,
				CvUrl:          "https://github.com/" + tableData.GithubId + "/" + tableData.CvRepoName,
			},
		)

		if ErrRowsScan != nil {
			appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: ErrRowsScan}
			return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", appErr)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		appErr := &domain.BackendError{Msg: "Members Not Found"}
		return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", appErr)
	}

	return ms, nil
}
