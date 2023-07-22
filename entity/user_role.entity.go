package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	Id        uuid.UUID `gorm:"primaryKey;column:user_role_id;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	Role      string `gorm:"column:role;type:varchar(10)"`
	UserId    uuid.UUID
}

func (UserRole) TableName() string {
	return "tb_user_role"
}
