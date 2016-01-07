package db

import (
	"database/sql"
	"fmt"

	"github.com/gregory90/go-webutils/try"

	"github.com/go-sql-driver/mysql"
)

func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	err = try.Do(func(attempt int) (bool, error) {
		var err interface{}
		err = txFunc(tx)
		if err != nil {
			switch v := err.(type) {
			case mysql.MySQLError:
				if err.(*mysql.MySQLError).Number == 1213 {
					return attempt < 5, err
				}
			default:
				return false, err
			}
		}
		return false, err
	})
	return err
}
