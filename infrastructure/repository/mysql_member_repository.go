package repository

import (
	"database/sql"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"golang.org/x/xerrors"
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
		ErrBackend := &domain.BackendError{Message: "Db.Prepare Error", Err: ErrPrepare}
		return nil, xerrors.Errorf("MysqlMemberRepository.Find: %w", ErrBackend)
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
			ErrBackend := &domain.BackendError{Message: "Member Not Found"}
			return nil, xerrors.Errorf("MysqlMemberRepository.Find: %w", ErrBackend)
		}

		ErrBackend := &domain.BackendError{Message: "rows.Scan Error", Err: ErrQuery}
		return nil, xerrors.Errorf("MysqlMemberRepository.Find: %w", ErrBackend)
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
		ErrBackend := &domain.BackendError{Message: "Db.Prepare Error", Err: ErrPrepare}
		return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", ErrBackend)
	}

	defer func() {
		ErrStmtClose := stmt.Close()
		if ErrStmtClose != nil {
			r.Logger.Error("stmt.Close() Fatal.", zap.Error(ErrStmtClose))
		}
	}()

	rows, ErrQuery := stmt.Query()

	if ErrQuery != nil {
		ErrBackend := &domain.BackendError{Message: "stmt.Query Error", Err: ErrQuery}
		return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", ErrBackend)
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
			ErrBackend := &domain.BackendError{Message: "rows.Scan Error", Err: ErrRowsScan}
			return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", ErrBackend)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.Id == 0 {
		ErrBackend := &domain.BackendError{Message: "Members Not Found"}
		return nil, xerrors.Errorf("MysqlMemberRepository.FindAll: %w", ErrBackend)
	}

	return members, nil
}
