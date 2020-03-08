-- 'portfolio_backend_test' というデータベースを作成
-- 'nekochans_test' というユーザー名のユーザーを作成
-- データベース 'portfolio_backend_test' への権限を付与
CREATE DATABASE IF NOT EXISTS portfolio_backend_test CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
GRANT ALL on portfolio_backend_test.* TO `nekochans_test`@`%` identified BY 'nekochans(Password2222)';
