package models

import (
	"time"
)

type Operators struct {
	Id           int       `gorm:"column:id;primary_key;NOT NULL;comment:'商户ID'"`
	Username     string    `gorm:"column:username;NOT NULL;comment:'账号'"`
	Avatar       string    `gorm:"column:avatar;default:;comment:'头像'"`
	Password     string    `gorm:"column:password;NOT NULL;comment:'密码'"`
	RoleName     string    `gorm:"column:role_name;default:NULL;comment:'角色名称\r\n'"`
	RoleId       int       `gorm:"column:role_id;default:0;NOT NULL;comment:'角色ID'"`
	LoginFailure uint8     `gorm:"column:login_failure;default:0;NOT NULL;comment:'失败次数'"`
	LoginIp      string    `gorm:"column:login_ip;default:NULL;comment:'登录IP'"`
	LoginTime    time.Time `gorm:"column:login_time;default:NULL;comment:'登录时间'"`
	Status       uint8     `gorm:"column:status;default:1;NOT NULL;comment:'账号状态，1：正常，2：禁用'"`
	UpdatedAt    uint64    `gorm:"column:updated_at;default:NULL;comment:'更新时间'"`
	DeletedAt    uint64    `gorm:"column:deleted_at;default:NULL;comment:'删除时间'"`
	CreatedAt    uint64    `gorm:"column:created_at;default:NULL;comment:'创建时间'"`
}

func (o *Operators) TableName() string {
	return "operators"
}
