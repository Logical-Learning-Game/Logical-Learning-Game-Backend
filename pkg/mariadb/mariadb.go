package mariadb

import (
	"database/sql"
	"llg_backend/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

func New(url string) (*sql.DB, error) {
	logger.GlobalLog.Info("connecting to mariadb...")
	conn, err := sql.Open("mysql", url)
	if err != nil {
		logger.GlobalLog.Debugw("connect to mariadb failed", "err", err)
	}

	logger.GlobalLog.Info("connect to mariadb successful")
	return conn, err
}
