package public

import (
	"bcw/server"
	"github.com/casbin/casbin/v2"
	"github.com/spf13/cast"
	"log"
)

type Permissions struct {
}

type RolePermission struct {
	RoleId int
	Path   string
}

// 获取角色权限
func (p Permissions) GetRolePermissions(roleId int) ([]RolePermission, error) {
	var rolePermissions []RolePermission

	db := server.HttpServers.DB.Table("role_permission").
		Joins("JOIN roles ON roles.id = role_permission.role_id").
		Joins("JOIN permissions ON permissions.id = role_permission.permission_id").
		Where("permissions.status = ?", 1).
		Select("roles.id as role_id, permissions.path")

	if roleId > 0 {
		db = db.Where("roles.id = ?", roleId)
	}
	err := db.Find(&rolePermissions).Error

	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

// 同步权限
func (p Permissions) SyncCasbinRules(e *casbin.Enforcer) {

	permissions, err := p.GetRolePermissions(0)
	if err != nil {
		log.Println(err)
		return
	}

	// 清除现有规则
	e.ClearPolicy()
	if err := e.SavePolicy(); err != nil {
		log.Panicln("error save policy: ", err)
	}

	//添加新规则
	for _, item := range permissions {
		_, _ = e.AddPolicy(cast.ToString(item.RoleId), item.Path, "*")
	}

}
