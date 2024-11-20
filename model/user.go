package model

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
    Email     string `gorm:"type:varchar(200)" json:"email"`
    Iphone    string `gorm:"type:varchar(20)" json:"iphone"`
}

// 增

// 删

// 改

// 查