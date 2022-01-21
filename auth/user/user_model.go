package user

import (
	"github.com/curatorc/cngf/database"
	"github.com/curatorc/cngf/hash"
	"time"
)

// User 用户模型
type User struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Phone    string `json:"-"`
	Password string `json:"-"`

	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	DeletedAt time.Time `gorm:"column:deleted_at;index;" json:"deleted_at,omitempty"`
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
