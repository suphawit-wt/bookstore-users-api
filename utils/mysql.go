package utils

import (
	"bookstore/users/utils/errors"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func MySQLErrorValidate(err error) *errors.RestErr {
	if err == sql.ErrNoRows {
		return errors.NewNotFoundError("Record Not Found.")
	}

	fmt.Println(err)

	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1062 {
			if strings.Contains(mysqlErr.Message, "users.email") {
				return errors.NewConflictError("User Email must be unique.")
			}
		}
	}

	return errors.NewInternalServerError("Error occur on database operation.")
}
