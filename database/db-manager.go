package database

import "github.com/BrandonEchols/common-go-utils/common"

type (
	IDbManager interface {
	}

	dbManager struct {
	}
)

func NewDb(cfg common.IConfigGetter) (IDbManager, error) {
	//TODO Connect to db
	return &dbManager{}, nil
}
