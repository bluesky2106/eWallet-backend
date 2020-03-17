package mysql

import (
	errs "github.com/bluesky2106/eWallet-backend/errors"
)

func emptyDBError(msg string) error {
	err := errs.New(errs.ECMySQLDBEmpty)
	return errs.WithMessage(err, msg)
}
