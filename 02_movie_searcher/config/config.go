package config

import (
	"net/http"
	"time"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Config is used to contains app's configurations
type Config struct {
	APIKeyOMDB string
	HTTPCli    *http.Client
	MySQLCli   *sqlx.DB
}

// New is used to get Config
func New() Config {
	// Initialize MySQL Client
	// TODO: move dataSourceName to file config
	mysqlCli := GetMySQLCli("root:pass@/i_view_stockbit_omdb")

	// Initialize HTTP Client
	httpCli := &http.Client{
		Timeout: 10 * time.Second,
	}

	return Config{
		APIKeyOMDB: "faf7e5bb&s", // TODO: should move this into file or
		HTTPCli:    httpCli,
		MySQLCli:   mysqlCli,
	}
}

func GetMySQLCli(dataSourceName string) *sqlx.DB {
	mysqlCli, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		logrus.Panicln(err)
	}
	return mysqlCli
}
