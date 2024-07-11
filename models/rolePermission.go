package models

type RolePermission struct {
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

func (r *RolePermission) TableName() string {
	return "role_permission"
}
