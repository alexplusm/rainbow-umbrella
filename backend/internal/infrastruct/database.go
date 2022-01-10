package infrastruct

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// https://habr.com/ru/company/oleg-bunin/blog/583558/

func NewDBConn() (*sql.DB, error) {
	dbCredentials := "root:example@tcp(172.22.0.3:3306)/rainbow-umbrella"
	db, err := sql.Open("mysql", dbCredentials)
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
