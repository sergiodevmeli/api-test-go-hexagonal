package mysql

import (
	"api_test_hexagonal/config"
	"database/sql"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", config.ConnectionUriMysql)
}
