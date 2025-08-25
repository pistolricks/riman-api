package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Orders      OrderModel
	Permissions PermissionModel
	Tokens      TokenModel
	Users       UserModel
	Clients     ClientModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Orders:      OrderModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
		Clients:     ClientModel{DB: db},
	}
}
