package utils

import (
	"bookstore/users/errs"
	"database/sql"
	"strings"

	mssql "github.com/microsoft/go-mssqldb"
)

func SQLServerErrorValidate(err error) error {
	if err == sql.ErrNoRows {
		return errs.NewNotFoundError("Record Not Found.")
	}

	if mssqlErr, ok := err.(mssql.Error); ok {
		if mssqlErr.Number == 2601 {
			if strings.Contains(mssqlErr.Message, "email") {
				return errs.NewConflictError("Email must be unique.")
			}
		}
	}

	return errs.NewUnexpectedError()
}
