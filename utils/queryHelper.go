package utils

import "database/sql"

func ExecQuery(db *sql.DB, query string, args ...any) (sql.Result, error) {
	sqlstmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer sqlstmt.Close()
	res, err := sqlstmt.Exec(args...)
	return res, err
}
