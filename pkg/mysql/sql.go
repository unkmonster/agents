package mysql

import "github.com/go-sql-driver/mysql"

func IsDuplicateEntryError(err error) bool {
	if myerr, ok := err.(*mysql.MySQLError); ok {
		return myerr.Number == 1062
	}
	return false
}
