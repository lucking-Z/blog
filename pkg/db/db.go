package db

import (
	"blogs/pkg/filter"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type ConfigDB struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

var instance *gorm.DB

func GetInstance() *gorm.DB {
	return instance
}

func transformClause(ex *filter.Expression) (clause.Clause, error) {
	switch ex.Relational {
	case filter.RelationalEq:
		return clause.Eq{Column: ex.AttrName, } fmt.Sprintf("%s = ?", ex.AttrName), nil
	case filter.RelationalNotEq:
		return fmt.Sprintf("%s != ?", ex.AttrName), nil
	case filter.RelationalGt:
		return fmt.Sprintf("%s > ?", ex.AttrName), nil
	case filter.RelationalGte:
		return fmt.Sprintf("%s >= ?", ex.AttrName), nil
	case filter.RelationalLt:
		return fmt.Sprintf("%s < ?", ex.AttrName), nil
	case filter.RelationalLte:
		return fmt.Sprintf("%s <= ?", ex.AttrName), nil
	case filter.RelationalIn:
		return fmt.Sprintf("%s in ?", ex.AttrName), nil
	default:
		return "", fmt.Errorf("expression relational operator is unknown:%d", ex.Relational)
	}
}

func Build(db *gorm.DB, ex *filter.Expression) error {
	query, err := transformQuery(ex)
	if err != nil {
		return err
	}
	db.Clauses(clause.Clause{Expression: clause.Eq{Column: true}})
	db.Where(query, ex.Val)
	sub := ex.GetSubExpression()
	for k, v := range sub {
		query, err := transformQuery(&sub[k])
		if err != nil {
			return err
		}
		if v.Logical == filter.LogicalOr {
			db.Or(query, v.AttrName)
		} else {
			db.Where(query, v.AttrName)
		}
		vSub := v.GetSubExpression()
		for k := range vSub {
			err = Build(db, &vSub[k])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InitDB(confObj ConfigDB) {
	sync.OnceFunc(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&&interpolateParams=true&loc=Local",
			confObj.Username,
			confObj.Password,
			confObj.Host,
			confObj.Port,
			confObj.DbName)
		loggerIns := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		)
		mysqlConfig := mysql.Config{DSN: dsn}
		gormConfig := &gorm.Config{
			Logger: loggerIns,
		}
		var err error
		instance, err = gorm.Open(mysql.New(mysqlConfig), gormConfig)
		if err != nil {
			panic(fmt.Sprintf("db init failed:%s", err.Error()))
		}

		sqlDB, err := instance.DB()
		if err != nil {
			panic(fmt.Sprintf("db init failed:%s", err.Error()))
		}
		//最大空闲连接数
		sqlDB.SetMaxIdleConns(10)
		//最大打开数量
		sqlDB.SetMaxOpenConns(100)
		//最大可服务时间
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
}
