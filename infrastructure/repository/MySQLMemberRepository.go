package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/domain"
	"log"
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
	// TODO 適切なエラー処理を行う
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	// TODO 適切なエラー処理を行う
	if err != nil {
		log.Fatal(err)
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

		// TODO 適切なエラー処理を行う
		// TODO 値がnilの際のパターンを考慮する
		if err != nil {
			log.Fatal(err)
		}
	}

	return ms, nil
}
