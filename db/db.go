package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/makarski/mybanana/log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"

	"github.com/makarski/mybanana/db/banana"
)

const maxConnectAttempts = 5

// NewDB returns
func NewDB(dsn string, maxOpen, maxIdle, attempt int) (*gorp.DbMap, error) {
	errfmt := "NewDB: %s"

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf(errfmt, err)
	}

	if err := conn.Ping(); err != nil {
		if attempt < maxConnectAttempts {
			log.Error.Printf(errfmt, err)
			time.Sleep(time.Second)
			return NewDB(dsn, maxOpen, maxIdle, attempt+1)
		}

		return nil, fmt.Errorf(errfmt, err)
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
