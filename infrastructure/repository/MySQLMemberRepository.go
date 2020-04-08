package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/domain"
	"golang.org/x/xerrors"
)

type MySQLMemberRepository struct {
	DB *sql.DB
}

type FindTableData struct {
	MemberID   int
	GitHubID   string
	AvatarURL  string
	CVRepoName string
}

func (m *MySQLMemberRepository) Find(memberID int) (*domain.Member, error) {
	sql := `
		SELECT
		  m.id AS MemberID,
		  mgu.github_id AS GitHubID,
		  mgu.avatar_url AS AvatarURL,
		  mgu.cv_repo_name AS CVRepoName
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

	err = stmt.QueryRow(memberID).Scan(&tableData.MemberID, &tableData.GitHubID, &tableData.AvatarURL, &tableData.CVRepoName)
	if err != nil {
		appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: err}
		return nil, xerrors.Errorf("MySQLMemberRepository.Find: %w", appErr)
	}

	// この条件の時はデータが1件も存在しない
	if tableData.MemberID == 0 {
		appErr := &domain.BackendError{Msg: "Member Not Found"}
		return nil, xerrors.Errorf("MySQLMemberRepository.Find: %w", appErr)
	}

	me := &domain.Member{
		ID:             tableData.MemberID,
		GitHubUserName: tableData.GitHubID,
		GitHubPicture:  tableData.AvatarURL,
		CvURL:          "https://github.com/" + tableData.GitHubID + "/" + tableData.CVRepoName,
	}

	return me, nil
}

type FindAllTableData struct {
	MemberID   int
	GitHubID   string
	AvatarURL  string
	CVRepoName string
}

func (m *MySQLMemberRepository) FindAll() (domain.Members, error) {
	sql := `
		SELECT
		  m.id AS MemberID,
		  mgu.github_id AS GitHubID,
		  mgu.avatar_url AS AvatarURL,
		  mgu.cv_repo_name AS CVRepoName
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
		err := rows.Scan(&tableData.MemberID, &tableData.GitHubID, &tableData.AvatarURL, &tableData.CVRepoName)
		ms = append(
			ms,
			&domain.Member{
				ID:             tableData.MemberID,
				GitHubUserName: tableData.GitHubID,
				GitHubPicture:  tableData.AvatarURL,
				CvURL:          "https://github.com/" + tableData.GitHubID + "/" + tableData.CVRepoName,
			},
		)

		if err != nil {
			appErr := &domain.BackendError{Msg: "rows.Scan Error", Err: err}
			return nil, xerrors.Errorf("MySQLMemberRepository.FindAll: %w", appErr)
		}
	}

	// この条件の時はデータが1件も存在しない
	if tableData.MemberID == 0 {
		appErr := &domain.BackendError{Msg: "Members Not Found"}
		return nil, xerrors.Errorf("MySQLMemberRepository.FindAll: %w", appErr)
	}

	return ms, nil
}
