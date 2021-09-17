package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load(".env")

var (
	ConnectionUriMysql = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user_mysql"),
		os.Getenv("pass_mysql"),
		os.Getenv("host_mysql"),
		os.Getenv("port_mysql"),
		os.Getenv("db_name_mysql"),
	)
	ConnectionUriMongo = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s",
		os.Getenv("user_mongo"),
		os.Getenv("pass_mongo"),
		os.Getenv("cluster_mongo"),
		os.Getenv("db_name_mongo"),
	)
)

const AllowedCORSDomain = "http://localhost"