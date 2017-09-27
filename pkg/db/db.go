package db

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"

	"github.com/makarski/mybanana/pkg/db/banana"
	"github.com/makarski/mybanana/pkg/log"
)

const maxConnectAttempts = 5

// NewDB returns
func NewDB(dsn string, maxOpen, maxIdle, attempt int) (*gorp.DbMap, error) {
	errctx := "NewDB"

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, errctx)
	}

	if err := conn.Ping(); err != nil {
		if attempt < maxConnectAttempts {
			log.Error.Printf("%s: %s", errctx, err)
			time.Sleep(time.Second)
			return NewDB(dsn, maxOpen, maxIdle, attempt+1)
		}

		return nil, errors.Wrap(err, errctx)
	}

	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	dbMap := &gorp.DbMap{
		Db: conn,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	// set mappings
	dbMap.AddTableWithName(banana.Banana{}, "banana").SetKeys(true, "ID")

	return dbMap, nil
}
