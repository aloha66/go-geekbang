package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func QuerySomeThing() error {
	errMock := sql.ErrNoRows
	if errors.Is(errMock,sql.ErrNoRows) {
		return errors.Wrap(errMock, "herr is something error")
	}
	return nil
}
