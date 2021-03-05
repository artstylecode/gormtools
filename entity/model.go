package entity

import "time"

type supp_admin struct {
	Id         int       `gorm:"column:id"`
	AdminName  string    `gorm:"column:admin_name"`
	Password   string    `gorm:"column:password"`
	RoleId     int       `gorm:"column:role_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	Email      string    `gorm:"column:email"`
	LastIp     string    `gorm:"column:last_ip"`
	LastTime   time.Time `gorm:"column:last_time"`
	IsDel      int       `gorm:"column:is_del"`
	LoginToken string    `gorm:"column:login_token"`
	Status     int       `gorm:"column:status"`
}
