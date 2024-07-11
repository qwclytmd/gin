package server

import (
	"bcw/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var HttpServers Servers

type Servers struct {
	DB     *gorm.DB
	Redis  *redis.Client
	Casbin *casbin.Enforcer
}

type CasbinEnforcer struct {
	Adapter   *gormadapter.Adapter
	RbacModel model.Model
}

func Initialize() {
	HttpServers = Servers{}
	HttpServers.Redis = config.ConnectionRedisClient()
	HttpServers.DB = config.ConnectionMySQL(config.GetViper().GetString("database.dbname"))
	HttpServers.Casbin = config.InitializeCasbinEnforcer(HttpServers.DB)

}
