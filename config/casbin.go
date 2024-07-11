package config

import (
	"log"

	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
)

func InitializeCasbinEnforcer(db *gorm.DB) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}

	m, err := model.NewModelFromFile("rbac_model.conf")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	return e
}
