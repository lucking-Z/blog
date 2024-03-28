package muser

import (
	"blogs/utils/db"
	"blogs/utils/log"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Add(ctx *gin.Context, user *User) (uint, error) {
	if user.Name == "" || user.Account == "" || user.PassWord == "" {
		msg := "user Add params error"
		log.Info(ctx, msg, log.Fields{"params": user})
		return 0, errors.New(msg)
	}

	result := db.Instance.Create(user)

	if result.Error != nil {
		log.Error(ctx, "db error", log.Fields{"params": user})
		return 0, result.Error
	}

	return user.ID, nil
}

func GetByAccount(ctx *gin.Context, account string) (*User, error) {
	if account == "" {
		msg := "GetByAccount  params is empty"
		log.Info(ctx, msg, log.Fields{"params": account})
		return nil, errors.New(msg)
	}
	user := &User{}
	result := db.Instance.First(user, "account = ?", account)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		log.Info(ctx, "db error", log.Fields{"error": result.Error, "params": account})
		return nil, result.Error
	}

	return user, nil
}
