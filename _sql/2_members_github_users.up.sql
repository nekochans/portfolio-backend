-- github_accountsにすべきか迷ったが、https://api.github.com/users というAPIに載っている情報を格納するテーブルなのでこの名前にした
CREATE TABLE `members_github_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int(10) unsigned NOT NULL,
  `github_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'GitHubのアカウント名 https://github.com/ + でprofileが見れる .e.g. keitakn',
  `avatar_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'GitHubのアカウントに設定されている画像URL',
  `cv_repo_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '職務経歴書のGitHubリポジトリ名 .e.g. cv, resume',
  `lock_version` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_members_github_users_01` (`member_id`),
  CONSTRAINT `fk_members_github_users_01` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin ROW_FORMAT=DYNAMIC;
