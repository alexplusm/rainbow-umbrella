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
		return nil, fmt.Errorf("[NewDBConn][1]: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[NewDBConn][2]: %w", err)
	}

	if err = execSQLScriptFromFile("scripts/schema.sql", db); err != nil {
		return nil, fmt.Errorf("[NewDBConn][3]: %w", err)
	}

	return db, nil
}

func execSQLScriptFromFile(filePath string, db *sql.DB) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("[execSQLScriptFromFile][1]: %w", err)
	}

	rawScript := string(fileContent)

	scripts := make([]string, 0, 8)

	for _, scriptChunk := range strings.Split(rawScript, ";\n") {
		trimmedScript := strings.TrimSpace(scriptChunk)
		if trimmedScript != "" {
			scripts = append(scripts, trimmedScript)
		}
	}

	for _, script := range scripts {
		if _, err = db.Exec(script); err != nil {
			return fmt.Errorf("[execSQLScriptFromFile][2]: %v: %w", script, err)
		}
	}

	return nil
}
