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
	`

	stmt, err := m.DB.Prepare(sql)
	if err != nil {
		return nil, xerrors.Errorf("MySQLMemberRepository.FindAll DB.Prepare Error: %w", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, xerrors.Errorf("MySQLMemberRepository.FindAll stmt.Query Error: %w", err)
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
			return nil, xerrors.Errorf("MySQLMemberRepository.FindAll rows.Scan Error: %w", err)
		}
	}

	return ms, nil
}
