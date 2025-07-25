package model

import (
	"time"
)

type UserDataModel struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`                         // 改为整型
	Username   string     `gorm:"column:username;type:varchar(255);not null" json:"username"` // 用户名
	Email      string     `gorm:"column:email;type:varchar(255);not null" json:"email"`       // 邮箱地址
	Gender     string     `gorm:"column:gender;type:varchar(16)" json:"gender"`               // 性别（如 male/female/other）
	Age        int        `gorm:"column:age;type:int" json:"age"`                             // 年龄
	CreateTime *time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`       // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`       // 更新时间
}

// TableName 设置表名为 `user`
func (UserDataModel) TableName() string {
	return "user"
}
