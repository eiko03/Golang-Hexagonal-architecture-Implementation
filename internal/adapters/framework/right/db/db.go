package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driveName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driveName, dataSourceName)
	if err != nil {
		log.Fatalf("DB conn failed: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da *Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("DB connection close failed: %v", err)
	}
}

func (da *Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").
		Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).
		ToSql()

	if err != nil {
		return err
	}
	_, err = da.db.Exec(queryString, args...)

	if err != nil {
		return err
	}
	return nil
}
