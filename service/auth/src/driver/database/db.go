package database

import (
	"github.com/srrrs-7/micro_service.git/auth/util/env"
	"gorm.io/gorm"
)

func NewDb(env *env.Config) *gorm.DB {
	tx := &gorm.DB{}
	return tx
}
