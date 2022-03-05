package infrastruct

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

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

	if err = execSQLScriptFromFile("scripts/schema.sql", db); err != nil {
		return nil, fmt.Errorf("[NewDBConn][3]: %+v", err)
	}

	return db, nil
}

func execSQLScriptFromFile(filePath string, db *sql.DB) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("[execSQLScriptFromFile][1]: %+v", err)
	}

	rawScript := string(fileContent)

	scripts := strings.Split(rawScript, ";\n")

	for _, script := range scripts {
		if _, err := db.Exec(script); err != nil {
			return fmt.Errorf("[execSQLScriptFromFile][2]: %+v", err)
		}
	}

	return nil
}
