package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
	IDUser uuid.UUID
	Username string
	Email string
	Password string
	Role string
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.IDUser = uuid.New()
	return
}