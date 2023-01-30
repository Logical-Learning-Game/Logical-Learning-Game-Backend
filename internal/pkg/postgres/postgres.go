package postgres

import (
	"database/sql"
	"llg_backend/pkg/logger"

	_ "github.com/lib/pq"
)

func New(uri string) (*sql.DB, error) {
	logger.GlobalLog.Infof("connecting to postgres at %s", uri)
	conn, err := sql.Open("postgres", uri)
	if err != nil {
		logger.GlobalLog.Errorw("connect to postgres failed", "err", err)
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		logger.GlobalLog.Errorw("verify postgres connection failed", "err", err)
		return nil, err
	}

	logger.GlobalLog.Infof("connecting to postgres successful")
	return conn, nil
}
