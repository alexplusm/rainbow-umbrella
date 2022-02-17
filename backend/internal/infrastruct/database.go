package infrastruct

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// https://habr.com/ru/company/oleg-bunin/blog/583558/

func NewDBConn(config *DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.URL)
	if err != nil {
		return nil, fmt.Errorf("[NewDBConn][1]: %+v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[NewDBConn][2]: %+v", err)
	}

	if err = initSchema(db); err != nil {
		return nil, fmt.Errorf("[NewDBConn][3]: %+v", err)
	}

	return db, nil
}

func initSchema(db *sql.DB) error {
	//db.Query()
	return nil
}
