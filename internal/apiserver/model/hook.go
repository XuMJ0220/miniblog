package model

import (
	"miniblog/internal/pkg/rid"

	"gorm.io/gorm"
)

var (
	UserPrefix = "user"
	PostPrefix = "post"
)

// AfterCreate 在创建数据库记录之后生成 userID.
func (m *UserM) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.NewResourceID(UserPrefix).New(uint64(m.ID))
	return tx.Save(m).Error
}

// AfterCreate 在创建数据库记录之后生成 postID.
func (m *PostM) AfterCreate(tx *gorm.DB) error {
	m.PostID = string(rid.NewResourceID(PostPrefix).New(uint64(m.ID)))
	return tx.Save(m).Error
}
